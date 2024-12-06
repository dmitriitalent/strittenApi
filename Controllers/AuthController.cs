using api.Entities;
using api.Dtos;
using api.Services;
using api.Exceptions;
using Microsoft.AspNetCore.Mvc;
using System.IdentityModel.Tokens.Jwt;
using api.Database;

namespace api.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class AuthController : ControllerBase
    {
        DatabaseContext db;
        RegistrationService RegistrationService;
        TokenService TokenService;
        LoginService LoginService;
        public AuthController(DatabaseContext db)
        {
            this.db = db;
            this.RegistrationService = new RegistrationService(this.db);
            this.TokenService = new TokenService(this.db);
            this.LoginService = new LoginService(this.db);
        }

        [HttpPost]
        [Route("registration")]
        public IActionResult Registration(RegistrationDTO registrationDTO)
        {
            // Validation
            string? error = RegistrationService.Validate(registrationDTO);
            Console.WriteLine(registrationDTO.Email);
            if (error != null) { throw new ValidationException<AuthController>(error); }

            // Checks if the login is exist in database
            error = RegistrationService.CheckLogin(registrationDTO.Login);
            if (error != null) {
                Console.WriteLine(registrationDTO.Login);
                throw new ValidationException<AuthController>(error); }

            // Add new User to database
            error = RegistrationService.AddUserToDatabase(registrationDTO);
            if (error != null) { throw new InternalServerException<AuthController>(error); }

            // Generate tokens: List of { AccessToken, RefreshToken }
            User user = db.Users.FirstOrDefault(user => user.Login == registrationDTO.Login);
            IList<JwtSecurityToken> tokens = TokenService.GenerateTokens(user);

            // Save RefreshToken to database
            TokenService.SaveToken(new JwtSecurityTokenHandler().WriteToken(tokens[1]), user);

            // Add tokens to response Cookies
            Response.Cookies.Append("AccessToken", new JwtSecurityTokenHandler().WriteToken(tokens[0]));
            Response.Cookies.Append("RefreshToken", new JwtSecurityTokenHandler().WriteToken(tokens[1]),
                new CookieOptions() { SameSite = SameSiteMode.None, HttpOnly = true, Secure = true } // Cookies is not available in JS of we use HttpOnly flag
            );

            return Ok(new JwtSecurityTokenHandler().WriteToken(tokens[0]));
        }

        [HttpPost]
        [Route("login")]
        public IActionResult Login(LoginDTO loginDTO)
        {
            // Finding user who have same login and password
            string error = LoginService.CheckUser(loginDTO);
            if (error != null)
                throw new BadRequestException<AuthController>(error);
            User user = db.Users.FirstOrDefault(user => user.Login == loginDTO.Login);

            // Generate tokens: List of { AccessToken, RefreshToken }
            IList<JwtSecurityToken> tokens = TokenService.GenerateTokens(user);

            // Save RefreshToken to database
            TokenService.SaveToken(new JwtSecurityTokenHandler().WriteToken(tokens[1]), user);

            // Add tokens to response Cookies
            Response.Cookies.Append("AccessToken", new JwtSecurityTokenHandler().WriteToken(tokens[0]));
            Response.Cookies.Append("RefreshToken", new JwtSecurityTokenHandler().WriteToken(tokens[1]),
                new CookieOptions() { SameSite = SameSiteMode.None, HttpOnly = true, Secure = true } // Cookies is not available in JS of we use HttpOnly flag
            );

            return Ok(new JwtSecurityTokenHandler().WriteToken(tokens[0]));
        }

        [HttpPost]
        [Route("logout")]
        public IActionResult Logout()
        {
            // Detele refreshToken from database
            string refreshToken;
            if (Request.Cookies.TryGetValue("RefreshToken", out refreshToken))
                TokenService.RemoveToken(refreshToken);

            // Clean cookies 
            Response.Cookies.Delete("RefreshToken");
            Response.Cookies.Delete("AccessToken");

            return Ok();
        }

        [HttpPost]
        [Route("refresh")]
        public IActionResult Refresh()
        {
            // Get RefreshToken from cookies
            string refreshToken;
            if (!Request.Cookies.TryGetValue("RefreshToken", out refreshToken))
                throw new UnauthorizedException<AuthController>("Ошибка авторизации #1");

            // Verify token`s VERIFY SIGNATURE 
            if (!TokenService.VerifyToken(refreshToken))
                throw new UnauthorizedException<AuthController>("Ошибка авторизации #2");

            // Check if exist RefreshToken in database
            if (!TokenService.ExistDbRefreshToken(refreshToken))
                throw new UnauthorizedException<AuthController>("Ошибка авторизации #3");

            // Get user to generate and save token 
            User user = TokenService.GetUserByToken(refreshToken);

            // Generate List of tokens { AccessToken, RefreshToken }
            IList<JwtSecurityToken> tokens = TokenService.GenerateTokens(user);

            // Save token to database
            TokenService.SaveToken(new JwtSecurityTokenHandler().WriteToken(tokens[1]), user);

            // Delete old cookies tokens
            Response.Cookies.Delete("RefreshToken");
            Response.Cookies.Delete("AccessToken");

            // Add tokens to response Cookies
            Response.Cookies.Append("AccessToken", new JwtSecurityTokenHandler().WriteToken(tokens[0]));
            Response.Cookies.Append("RefreshToken", new JwtSecurityTokenHandler().WriteToken(tokens[1]),
                new CookieOptions() { SameSite = SameSiteMode.None, HttpOnly = true, Secure = true } // Cookies is not available in JS of we use HttpOnly flag
            );

            return Ok(new JwtSecurityTokenHandler().WriteToken(tokens[0]));
        }
    }
}
