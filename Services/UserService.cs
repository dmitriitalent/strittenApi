using Newtonsoft.Json;
using api.Entities;
using api.Database;
using api.Exceptions;
using Microsoft.EntityFrameworkCore;
using api.Dtos;
using System.Buffers.Text;

namespace api.Services
{

    public class UserService
    {
        DatabaseContext db;
        ValidationService ValidationService;
        public UserService(DatabaseContext db)
        {
            this.db = db;
            this.ValidationService = new ValidationService(db);
        }
        
        public User EditUser(EditUserDTO UserDTO)
        {
            User User = db.Users
                .Include(user => user.Profile)
                .Include(user => user.OrganizedEvents)
                .Include(user => user.VisitedEvents)
                .FirstOrDefault(user => user.Id == UserDTO.UserId);

     
            try
            {
                ValidationService.ValidateLength(UserDTO.Surname, 2);
            }
            catch (ValidationException<ValidationService> ex)
            {
                throw new ValidationException<UserService>("Surname must be longer");
            }
            try
            {
                ValidationService.ValidateLength(UserDTO.Name, 2);
            }
            catch (ValidationException<ValidationService> ex)
            {
                throw new ValidationException<UserService>("Name must be longer");
            }
            if (UserDTO.Login != User.Login)
                ValidationService.ValidateLogin(UserDTO.Login);
            ValidationService.ValidateEmail(UserDTO.Email);
            
            User.Login = UserDTO.Login;
            User.Profile.Name = UserDTO.Name;
            User.Profile.Surname = UserDTO.Surname;
            User.Profile.Email = UserDTO.Email;

            db.SaveChanges();

            return User;
        }
    }
}
