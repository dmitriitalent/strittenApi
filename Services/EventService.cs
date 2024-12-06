using Newtonsoft.Json;
using api.Entities;
using api.Database;
using api.Exceptions;
using Microsoft.EntityFrameworkCore;
using api.Dtos;
using System.Buffers.Text;

namespace api.Services
{
	public class EventService
	{
		DatabaseContext db;
		public EventService(DatabaseContext db)
		{
			this.db = db;
		}
		public Event GetEvent(int Id)
		{

			Event? Event = db
				.Events
				.Include(Event => Event.General)
				.Include(Event => Event.Additionals)
				.Include(Event => Event.Visiters)
				.ThenInclude(v => v.OrganizedEvents)
				.Include(Event => Event.Organizer)
				.ThenInclude(User => User.Profile)
				.FirstOrDefault(e => e.Id == Id);
			if (Event == null)
			{
				throw new NotFoundException<EventService>($"Event with id \"{Id}\" not found.");
			}
			
			return Event;
		}
		public Event GetEventByInviteCode(string InviteCode)
		{
			Event? Event = db.Events.Include(e => e.General).FirstOrDefault(e => e.General.InviteCode == InviteCode);
			if(Event == null)
			{
				throw new NotFoundException<EventService>("Incorrect invite code.");
			}
			return Event;
		}

		public IList<Event> GetEvents(int Count, int Page)
		{
			IList<Event> result = new List<Event>();
			IList<Event> events = db
				.Events
				.Include(Event => Event.General)
				.Include(Event => Event.Additionals)
				.Include(Event => Event.Visiters)
				.Include(Event => Event.Organizer)
				.ToList();

			if (events.Count < Page * Count)
			{
				var lastEvents = events.TakeLast(Count);
				foreach (var e in lastEvents)
				{
					result.Add(e);
				}
			}
			else
			{
				for (int i = (Page - 1) * Count; i < Page * Count; i++)
				{
					result.Add(events[i]);
				}
			}

			return result;
		}

		public Event CreateEvent(CreateEventDTO EventData, int OgranizerId)
		{

			Event NewEvent = new Event();
			NewEvent.OrganizerId = OgranizerId;
			User? Organizer = db.Users.FirstOrDefault(user => user.Id == EventData.OrganizerId);
			if(Organizer == null)
			{
				throw new api.Exceptions.UnauthorizedException<EventService>();
			}
			NewEvent.Visiters.Add(
					Organizer
				);

			IList<EventAdditional> NewEventAdditionals = new List<EventAdditional>();
			foreach (var additional in EventData.Additionals)
			{
				NewEventAdditionals.Add(new EventAdditional()
				{
					Key = additional.Key,
					Value = additional.Value,
					Event = NewEvent
				});
			}

			EventGeneral NewEventGeneral = new EventGeneral()
			{
				Date = EventData.General.Date,
				Description = EventData.General.Description,
				Fundraising = EventData.General.Fundraising,
				HideDate = EventData.General.HideDate,
				HidePlace = EventData.General.HidePlace,
				Name = EventData.General.Name,
				Place = EventData.General.Place,
				Private = EventData.General.Private,
				Event = NewEvent,
				HideInviteCode = EventData.General.HideInviteCode,
				InviteCode = ""
			};

			NewEvent.Additionals = NewEventAdditionals;
			NewEvent.General = NewEventGeneral;

			db.Events.Add(NewEvent);
			db.SaveChanges();

			Int32 unixTime = (int)DateTime.UtcNow.Subtract(new DateTime(1970, 1, 1)).TotalSeconds;
			Random rnd = new Random(unixTime);
			NewEvent.General.InviteCode = SecurityService.Base64Encode(
				(rnd.Next(int.MaxValue)).ToString()
			);
			db.SaveChanges();

			return NewEvent;
		}
		
		public Event DeleteEvent(int EventId)
		{
			Event Event = db.Events.FirstOrDefault(e => e.Id == EventId)!;
			db.Events.Remove(Event);
			db.SaveChanges();
			return Event;
		}

		public Event EditEvent(EditEventDTO EventData, int UserId)
		{
			Event Event = GetEvent(EventData.Id);

			Event.General.Private = EventData.General.Private;
			Event.General.Date = EventData.General.Date;
			Event.General.Description = EventData.General.Description;
			Event.General.Fundraising = EventData.General.Fundraising;
			Event.General.Name = EventData.General.Name;
			Event.General.Place = EventData.General.Place;
			Event.General.HidePlace = EventData.General.HidePlace;
			Event.General.HideDate = EventData.General.HideDate;

			Event.Additionals.Clear();
			foreach (var additional in EventData.Additionals)
			{
				Event.Additionals.Add(new EventAdditional()
				{
					Key = additional.Key,
					Value = additional.Value,
				});
			}

			db.SaveChanges();

			return Event;
		}

		public IList<Event> GetVisitedEvents(int UserId)
		{
			return
				db
				.Users
				.Include(user => user.VisitedEvents)
				.ThenInclude(e => e.General)
				.Include(user => user.VisitedEvents)
				.ThenInclude(e => e.Additionals)
				.FirstOrDefault(e => e.Id == UserId)
				.VisitedEvents;
		}

		public Event AddVisiter(int UserId, int EventId)
		{
			Event? Event = db.Events.Include(e => e.Visiters).FirstOrDefault(e => e.Id == EventId);
			if (Event == null)
			{
				throw new NotFoundException<EventService>($"Event with id \"{EventId}\" not found.");
			}
			User? User = db.Users.FirstOrDefault(user => user.Id == UserId);
			if (User == null)
			{
				throw new NotFoundException<EventService>($"User with id \"{UserId}\" not found.");
			}
			if (Event.Visiters.FirstOrDefault(user => user.Id == UserId) != null)
			{
				throw new BadRequestException<EventService>("User is already registered for the event");
			}
			Event.Visiters.Add(User);
			db.SaveChanges();
			return Event;
		}
		public Event RemoveVisiter(int UserId, int EventId)
		{
			Event? Event = db.Events.Include(e => e.Visiters).FirstOrDefault(e => e.Id == EventId);
			if (Event == null)
			{
				throw new NotFoundException<EventService>($"Event with id \"{EventId}\" not found.");
			}
			User? User = db.Users.FirstOrDefault(user => user.Id == UserId);
			if (User == null)
			{
				throw new NotFoundException<EventService>($"User with id \"{UserId}\" not found.");
			}
			if (Event.Visiters.FirstOrDefault(user => user.Id == UserId) == null)
			{
				throw new BadRequestException<EventService>("User is not registered for the event");
			}
			Event.Visiters.Remove(User);
			db.SaveChanges();
			return Event;
		}
	} 
}
