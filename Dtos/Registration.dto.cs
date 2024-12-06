namespace api.Dtos
{
	public class RegistrationDTO
	{
        public string Login { get; set; }
        public string Password { get; set; }
        public string PasswordConfirmed { get; set; }
        public string Email { get; set; }
        public string Surname { get; set; }
        public string Name { get; set; }
	}
}
