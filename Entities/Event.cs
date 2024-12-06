namespace api.Entities
{
    public class Event
    {
        public int Id { get; set; }
        public EventGeneral General { get; set; }
        public IList<EventAdditional> Additionals { get; set; } = new List<EventAdditional>();
        public User Organizer { get; set; }
        public int OrganizerId { get; set; }
        public IList<User> Visiters { get; set; } = new List<User>();
    }
}
