using api.Entities;
using api.Database;
using api.Dtos;

namespace api.Services;

public class RegistrationService
{
	DatabaseContext db;
	ValidationService ValidationService;
	HashPasswordService HashPasswordService;
	public RegistrationService(DatabaseContext db)
	{
		this.db = db;
		this.ValidationService = new ValidationService(db);
		this.HashPasswordService = new HashPasswordService();
	}

	public string? Validate(RegistrationDTO registrationDTO)
	{
		string error = null;
		error = ValidationService.ValidateEmail(registrationDTO.Email);
		if (error != null) 
			return error;
			
		error = ValidationService.ValidatePassword(registrationDTO.Password);
		if(error != null)
			return error;

		error = ValidationService.ValidatePasswordConfirm(registrationDTO.Password, registrationDTO.PasswordConfirmed);
		if(error != null)
			return error;

		error = ValidationService.ValidateLength(registrationDTO.Surname, 2);
		if (error != null)
			return "Введите фамилию";

		error = ValidationService.ValidateLength(registrationDTO.Name, 2);
		if (error != null)
			return "Введите имя";

		error = ValidationService.ValidateLogin(registrationDTO.Login);
		if (error != null)
			return "Логин " + error;

		return null;
	}

	public string? CheckLogin(string login)
	{
		User user = db.Users.FirstOrDefault(user => user.Login == login);
		if (user != null) { return "Логин уже используется"; }
		return null;
	}

	public string? AddUserToDatabase(RegistrationDTO registrationDTO)
	{
		User user = new User()
		{
			Login = registrationDTO.Login,
			Password = HashPasswordService.Make(registrationDTO.Password),
			Role = Roles.User
		};
		try { db.Users.Add(user); }
		catch (Exception ex)
		{
			return "Не удалось добавить пользователя в базу данных";
		}
		
		Profile userProfile = new Profile()
		{
			Email = registrationDTO.Email,
			Name = registrationDTO.Name,
			Surname = registrationDTO.Surname,
			User = user
		};
		try { db.Profiles.Add(userProfile); }
		catch (Exception ex)
		{
			return "Не удалось добавить пользователя в базу данных";
		}

		try { db.SaveChanges(); }
		catch (Exception ex)
		{
			return "Не удалось сохранить пользователя";
		}
		return null;
	}
}

static class Roles
{
	public static string Admin = "Admin";
	public static string Moderator = "Moderator";
	public static string User = "User";
}