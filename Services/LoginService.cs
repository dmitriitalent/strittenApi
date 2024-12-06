using api.Dtos;
using api.Entities;
using api.Database;
using System.IdentityModel.Tokens.Jwt;

namespace api.Services
{
	public class LoginService
	{
		DatabaseContext db;
		TokenService TokenService;
		private HashPasswordService HashPasswordService;

        public LoginService(DatabaseContext db) 
		{
			this.db = db;
            this.TokenService = new TokenService(db);
            this.HashPasswordService = new HashPasswordService();

		}

        public string? CheckUser(LoginDTO loginDTO)
        {
	        IQueryable<User> users = db.Users.Where(user => user.Login == loginDTO.Login);
	        var passwordCompared = false;
	        foreach (var user in users)
	        {
		        passwordCompared = HashPasswordService.Compare(user.Password, loginDTO.Password);
	        }
			if (!passwordCompared) 
			{
				return "Неверный логин или пароль";
			}
			return null;

			
		}

	}
}
