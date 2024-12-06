using api.Entities;

namespace api.Dtos
{
    public class OrganizedEventsDTO : List<EventDTO>
    {
        public OrganizedEventsDTO(IList<Event> OrganizedEvents, UserDTO parentOrganizer = null)
        {
            if (OrganizedEvents != null)
            {
                foreach (var Event in OrganizedEvents)
                {
                    this.Add(new EventDTO(Event, parentOrganizer: parentOrganizer));
                }
            }
        }
    }
}
