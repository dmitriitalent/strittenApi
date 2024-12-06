using api.Entities;
using Microsoft.IdentityModel.Tokens;

namespace api.Dtos
{
	public class EventDTO : Entities.Event
	{
		public EventDTO(Event Event, 
			EventGeneralDTO parentEventGeneral = null, 
			UserDTO parentOrganizer = null,
			UserDTO parentVisitor = null
        )
		{
            this.Id = Event.Id;
			this.Additionals = Event.Additionals;
			if(parentEventGeneral == null && Event.General != null)
				this.General = new EventGeneralDTO(Event.General, this);

			if (parentOrganizer == null && Event.Organizer != null)
			{
				this.Organizer = new UserDTO(Event.Organizer, this);
				this.OrganizerId = Event.OrganizerId;
            }
			
			if(parentVisitor == null)
			{
				this.Visiters = new EventVisitersDTO(Event.Visiters, parentEvent: this);
			}
		}

        public EventGeneralDTO? General { get; set; } = null;
		public EventVisitersDTO? Visiters { get; set; } = null;
		public UserDTO? Organizer { get; set; } = null;
    }
}
