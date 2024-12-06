namespace api.Entities
{
    public class EventAdditional
    {
        public int Id { get; set; }
        public string Key { get; set; }
        public string Value { get; set; }

        public Event Event { get; set; }
        public int EventId { get; set; }
    }
}
