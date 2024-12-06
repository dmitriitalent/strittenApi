using api.Entities;

namespace api.Dtos
{
	public class EventGeneralDTO : EventGeneral
	{
		public EventGeneralDTO(EventGeneral EventGeneral, EventDTO parent = null)
		{
			Id = EventGeneral.Id;
			Name = EventGeneral.Name;
			Description = EventGeneral.Description;
			InviteCode = EventGeneral.InviteCode;
			HideInviteCode = EventGeneral.HideInviteCode;
			Private = EventGeneral.Private;
			HideDate = EventGeneral.HideDate;
			HidePlace = EventGeneral.HidePlace;
			Fundraising = EventGeneral.Fundraising;
			Date = EventGeneral.Date;
			Place = EventGeneral.Place;

			if (parent == null)
				Event = new EventDTO(EventGeneral.Event, this);
			else
				Event = null;

			EventId = EventGeneral.EventId;
		}

		public string? InviteCode { get; set; } = null;
		public DateTime? Date { get; set; } = null;
		public string? Place { get; set; } = null;
		public EventDTO? Event { get; set; } = null;
	}
}
