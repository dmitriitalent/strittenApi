using api.Entities;

namespace api.Dtos
{
    public class VisitedEventsDTO : List<EventDTO>
    {
        public VisitedEventsDTO(IList<Event> VisitedEvents, UserDTO parentVisitor)
        {
            if (VisitedEvents != null)
            {
                foreach (var Event in VisitedEvents)
                {
                    this.Add(new EventDTO(Event, parentVisitor: parentVisitor));
                }
            }
        }
    }
}
