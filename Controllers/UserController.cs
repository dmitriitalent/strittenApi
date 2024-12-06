using api.Entities;
using api.Dtos;
using api.Services;
using api.Exceptions;
using Microsoft.AspNetCore.Mvc;
using System.IdentityModel.Tokens.Jwt;
using api.Database;
using Newtonsoft.Json;
using Microsoft.EntityFrameworkCore;
using Microsoft.IdentityModel.Tokens;

namespace api.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class UserController : ControllerBase
    {
        DatabaseContext db;
        EventService EventService;
        TokenService TokenService;
        UserService UserService;

        public UserController(DatabaseContext db)
        {
            this.db = db;
            this.EventService = new EventService(db);
            this.TokenService = new TokenService(db);
            this.UserService = new UserService(db);
        }

        [HttpGet]
        [Route("getProfileData")]
        public IActionResult GetProfileData(int Id)
        {
            return Ok(
                Newtonsoft.Json.JsonConvert.SerializeObject(
                    db.Profiles.FirstOrDefault(profile => profile.Id == Id)
                )
            );
        }
        [HttpGet]
        [Route("getUser")]
        public IActionResult Get(int Id)
        {
            User User = db.Users.Include(u => u.Profile).FirstOrDefault(user => user.Id == Id);
            if (User == null)
            {
                throw new NotFoundException<UserController>($"User with Id \"{Id}\" not found.");
            }

            return Ok(
                    JsonConvert.SerializeObject(
                        new UserDTO(User)
                    )
                );
        }

        [HttpGet]
        [Route("getVisitedEvents")]
        public IActionResult GetVisitedEvents(int Id)
        {
            IList<Event> VisitedEvents = EventService.GetVisitedEvents(Id);
            foreach (var Event in VisitedEvents)
            {
                Event.Organizer = null;
            }
            return Ok(
                Newtonsoft.Json.JsonConvert.SerializeObject(
                    VisitedEvents,
                    new JsonSerializerSettings()
                    {
                        ReferenceLoopHandling = ReferenceLoopHandling.Ignore
                    }
                )
            );
        }

        [HttpPost]
        [Route("edit")]
        public IActionResult EditUser(EditUserDTO UserDTO)
        {
            string RefreshToken;
            Request.Cookies.TryGetValue("RefreshToken", out RefreshToken);
            TokenService.VerifyToken(RefreshToken);

            User User = TokenService.GetUserByToken(RefreshToken);
            
            if(UserDTO.UserId != User.Id)
            {
                throw new ForbiddenException<UserController>();
            }

            User = UserService.EditUser(UserDTO);
            
            return Ok(JsonConvert.SerializeObject(new UserDTO(User)));
        }
    }
}
