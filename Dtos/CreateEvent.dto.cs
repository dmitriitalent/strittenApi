namespace api.Dtos
{
    public class CreateEventDTO
    {
        public CreateEventGeneralDTO General { get; set; }
        public List<CreateEventAdditionalDTO> Additionals { get; set; }
        public int OrganizerId { get; set; }
    }
    public class CreateEventGeneralDTO
    {
        public string Name { get; set; }
        public string Description { get; set; }
        public bool HideInviteCode { get; set; }
        public bool Private { get; set; }
        public bool HideDate { get; set; }
        public bool HidePlace { get; set; }
        public bool Fundraising { get; set; }
        public DateTime Date { get; set; }
        public string Place { get; set; }
    }
    public class CreateEventAdditionalDTO
    {
        public string Key { get; set; }
        public string Value { get; set; }
    }
}
