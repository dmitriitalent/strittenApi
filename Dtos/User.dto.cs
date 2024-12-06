using api.Entities;
using Microsoft.IdentityModel.Tokens;

namespace api.Dtos
{
	public class UserDTO : User
	{
		public UserDTO(User User, EventDTO parentEvent = null, ProfileDTO parentProfile = null)
		{

			this.Id = User.Id;
			this.Login = User.Login;
			this.Role = User.Role;
			if (parentProfile == null && User.Profile != null)
				this.Profile = new ProfileDTO(User.Profile, this);

			if(parentEvent == null)
			{
				this.OrganizedEvents = new OrganizedEventsDTO(User.OrganizedEvents, parentOrganizer: this);
			}
			
			if(parentEvent == null)
			{
				this.VisitedEvents = new VisitedEventsDTO(User.VisitedEvents, parentVisitor: this);
			}

			this.Password = null;
			this.RefreshToken = null;
		}

		public ProfileDTO? Profile { get; set; } = null;
		public OrganizedEventsDTO? OrganizedEvents { get; set; } = null;
		public VisitedEventsDTO? VisitedEvents { get; set; } = null;
	}
}
