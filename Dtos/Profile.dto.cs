using api.Entities;

namespace api.Dtos
{
	public class ProfileDTO : Profile
	{
		public ProfileDTO(Profile Profile, UserDTO parentUser) 
		{

			this.Id = Profile.Id;
			this.Name = Profile.Name;
			this.Surname = Profile.Surname;
			this.Email = Profile.Email;

			if (parentUser == null && Profile.User != null)
			{
				this.User = new UserDTO(Profile.User, parentProfile: this);
				this.UserId = Profile.UserId;
			}
			else
				this.User = null;
        }
		public UserDTO? User { get; set; } = null;

	}
}
