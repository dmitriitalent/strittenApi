using api.Database;
using api.Dtos;
using api.Entities;
using api.Services;
using api.Exceptions;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Newtonsoft.Json;
using Microsoft.AspNetCore.Http.HttpResults;
using Microsoft.IdentityModel.Tokens;

namespace api.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class EventController : ControllerBase
    {
        DatabaseContext db;
        EventService EventService;
        TokenService TokenService;
        public EventController(DatabaseContext db)
        {
            this.db = db;
            this.EventService = new EventService(db);
            this.TokenService = new TokenService(db);
        }

        [HttpPost]
        [Route("Create")]
        public IActionResult Create([FromBody] CreateEventDTO EventData)
        {
            string RefreshToken;
            Request.Cookies.TryGetValue("RefreshToken", out RefreshToken);
            TokenService.VerifyToken(RefreshToken);

            int OrganizerId = TokenService.GetUserByToken(RefreshToken).Id;

            Event NewEvent = EventService.CreateEvent(EventData, OrganizerId);

            return Ok(JsonConvert.SerializeObject(
                new EventDTO(NewEvent),
                new JsonSerializerSettings()
                {
                    ReferenceLoopHandling = ReferenceLoopHandling.Ignore,
                }
                ));
        }

        [HttpGet]
        [Route("GetEvent")]
        public IActionResult Get(int Id)
        {

            string RefreshToken;
            Request.Cookies.TryGetValue("RefreshToken", out RefreshToken);
            TokenService.VerifyToken(RefreshToken);

            int UserId = TokenService.GetUserByToken(RefreshToken).Id;

            Event Event = EventService.GetEvent(Id);
            if (Event.General.Private)
            {
                User? User = db.Users.Include(user => user.VisitedEvents).FirstOrDefault(user => user.Id == UserId);
                if (User == null)
                {
                    throw new NotFoundException<EventService>($"User with id \"{UserId}\" not found.");
                }
                bool UserIsVisiter = false;
                foreach (Event VisitedEvent in User.VisitedEvents)
                {
                    if (VisitedEvent.Id == Id)
                    {
                        UserIsVisiter = true;
                        break;
                    }

                }
                if (UserId != Event.OrganizerId && !UserIsVisiter)
                {
                    throw new ForbiddenException<EventService>("You are not visiter or organizer.");
                }
            }
            EventDTO EventDTO = new EventDTO(Event);
            if (UserId != Event.OrganizerId)
            {
                if (Event.General.HideDate == true)
                    EventDTO.General.Date = null;
                if (Event.General.HidePlace == true)
                    EventDTO.General.Place = null;
                if (Event.General.HideInviteCode == true)
                    EventDTO.General.InviteCode = null;
            }
            return Ok(
                JsonConvert.SerializeObject(
                    EventDTO,
                    new JsonSerializerSettings()
                    {
                        ReferenceLoopHandling = ReferenceLoopHandling.Ignore
                    }
                )
            );
        }
        [HttpGet]
        [Route("GetEvents")]
        public IActionResult GetEvents(int Count, int Page)
        {
            IList<Event> result = EventService.GetEvents(Count, Page);
            IList<EventDTO> EventsDTO = new List<EventDTO>();
            foreach(var Event in result)
            {
                EventsDTO.Add(new EventDTO(Event));
            }

            return Ok(
                JsonConvert.SerializeObject(
                    EventsDTO,
                    new JsonSerializerSettings()
                    {
                        ReferenceLoopHandling = ReferenceLoopHandling.Ignore
                    })
                );
        }

        [HttpPost]
        [Route("EditEvent")]
        public IActionResult EditEvent([FromBody] EditEventDTO EventData)
        {
            string RefreshToken;
            Request.Cookies.TryGetValue("RefreshToken", out RefreshToken);
            TokenService.VerifyToken(RefreshToken);

            User User = TokenService.GetUserByToken(RefreshToken);
            if (User.Id != EventData.OrganizerId)
            {
                throw new ForbiddenException<EventController>();
            }

            EventDTO EditedEvent = new EventDTO(EventService.EditEvent(EventData, User.Id));

            return Ok(
                JsonConvert.SerializeObject(
                    EditedEvent,
                    new JsonSerializerSettings()
                    {
                        ReferenceLoopHandling = ReferenceLoopHandling.Ignore
                    })
                );
        }
        [HttpPost]
        [Route("Delete")]
        public IActionResult DeleteEvent([FromBody] DeleteEventDTO EventData)
        {
            string RefreshToken;
            Request.Cookies.TryGetValue("RefreshToken", out RefreshToken);
            TokenService.VerifyToken(RefreshToken);
            
            Event Event = db.Events.Include(e => e.Organizer).FirstOrDefault(e => e.Id == EventData.Id);
            if (Event == null)
            {
                throw new NotFoundException<EventController>($"Event with id \"{EventData.Id}\" not found.");
            }
            if (TokenService.GetUserByToken(RefreshToken).Id != Event.OrganizerId)
            {
                throw new ForbiddenException<EventController>();
            }
            Event = EventService.DeleteEvent(EventData.Id);

            return Ok(
                JsonConvert.SerializeObject(
                    new EventDTO(Event),
                    new JsonSerializerSettings()
                    {
                        ReferenceLoopHandling = ReferenceLoopHandling.Ignore
                    })
                );
        }

        [HttpPost]
        [Route("AddVisiter")]
        public IActionResult AddVisiter([FromBody] AddVisiterDTO Dto)
        {
            Event Event = EventService.AddVisiter(Dto.UserId, Dto.EventId);

            return Ok(
                JsonConvert.SerializeObject(
                    new EventDTO(Event),
                    new JsonSerializerSettings()
                    {
                        ReferenceLoopHandling = ReferenceLoopHandling.Ignore
                    })
                );
        }
        [HttpPost]
        [Route("RemoveVisiter")]
        public IActionResult RemoveVisiter([FromBody] RemoveVisiterDTO Dto)
        {
            string RefreshToken;
            Request.Cookies.TryGetValue("RefreshToken", out RefreshToken);
            TokenService.VerifyToken(RefreshToken);
            Event? Event = db.Events.Include(e => e.Organizer).FirstOrDefault();
            if (Event == null)
            {
                throw new NotFoundException<EventController>($"Event with id \"{Dto.EventId}\" not found.");
            }
            if (
                TokenService.GetUserByToken(RefreshToken).Id != Event.OrganizerId &&
                TokenService.GetUserByToken(RefreshToken).Id != Dto.UserId
                )
            {
                throw new ForbiddenException<EventController>();
            }

            Event EditedEvent = EventService.RemoveVisiter(Dto.UserId, Dto.EventId);

            return Ok(
                JsonConvert.SerializeObject(
                    new EventDTO(EditedEvent),
                    new JsonSerializerSettings()
                    {
                        ReferenceLoopHandling = ReferenceLoopHandling.Ignore
                    })
                );
        }

        [HttpGet]
        [Route("getEventByInviteCode")]
        public IActionResult GetEventByInviteCode(string InviteCode)
        {
            string RefreshToken;
            Request.Cookies.TryGetValue("RefreshToken", out RefreshToken);
            TokenService.VerifyToken(RefreshToken);

            User User = TokenService.GetUserByToken(RefreshToken);

            Event Event = EventService.GetEventByInviteCode(InviteCode);
            Event = EventService.AddVisiter(User.Id, Event.Id);
            return Ok(
                JsonConvert.SerializeObject(
                    new EventDTO(Event),
                    new JsonSerializerSettings()
                    {
                        ReferenceLoopHandling = ReferenceLoopHandling.Ignore
                    })
                );
        }
    }
}
