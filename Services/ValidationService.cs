using Microsoft.Extensions.Diagnostics.HealthChecks;
using Microsoft.IdentityModel.Tokens;
using System.Text.RegularExpressions;
using api.Exceptions;
using api.Database;

namespace api.Services
{
    public class ValidationService
    {
        DatabaseContext db;
        public ValidationService(DatabaseContext db) 
        {
            this.db = db;
        }

        public string? ValidateLength(string word, int minLength = 0, double maxLength = double.PositiveInfinity)
        {
            if (word.Length < minLength)
            {
                throw new ValidationException<ValidationService>("shorter then " + minLength + " chars");
            }
            if (word.Length > maxLength)
            {
                throw new ValidationException<ValidationService>("longer then " + maxLength + " chars");
            }
            return null;
        }

        public string? ValidateLogin(string Login, int minLength = 2, double maxLength = 20)
        {
            if(db.Users.FirstOrDefault(user => user.Login == Login) != null)
            {
                throw new ValidationException<ValidationService>("Login is already in use");
            }

            if (Login.Length < minLength)
            {
                throw new ValidationException<ValidationService>($"The login must have a length from {minLength} to {maxLength}");
            }
            if (Login.Length > maxLength)
            {
                throw new ValidationException<ValidationService>($"The login must have a length from {minLength} to {maxLength}");
            }
            return null;
        }

        public string? ValidateEmail(string email)
        {
            Regex regex = new Regex(@"^([\w\.\-]+)@([\w\-]+)((\.(\w){2,3})+)$");
            Match match = regex.Match(email);
            if (!match.Success)
            {
                throw new ValidationException<ValidationService>("Enter your email");
            }
            else { return null; }
        }

        public string? ValidatePassword(string password)
        {
            if (password.Length < 7)
            {
                throw new ValidationException<ValidationService>("Password too short");
            }
            for (int i = 0; i < 10; i++)
            {
                if (password.Contains(i.ToString()))
                {
                    break;
                }
                if (i == 9)
                {
                    throw new ValidationException<ValidationService>("Password must contain a digit");
                }
            }
            return null;
        }

        public string? ValidatePasswordConfirm(string password, string passwordConfirm)
        {
            if (passwordConfirm != password)
                throw new ValidationException<ValidationService>("Passwords must be equal");
            return null;
        }
    }
}
