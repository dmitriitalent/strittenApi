namespace api.Entities
{
    public class EventGeneral
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public string InviteCode { get; set; }
        public bool HideInviteCode { get; set; }
        public bool Private { get; set; }
        public bool HideDate { get; set; }
        public bool HidePlace { get; set; }
        public bool Fundraising { get; set; }
        public DateTime Date { get; set; }
        public string Place { get; set; }

        public Event Event { get; set; }
        public int EventId { get; set; }
    }
}
