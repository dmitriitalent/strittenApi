using System.ComponentModel.DataAnnotations;

namespace api.Entities
{
    public class User
    {
        public int Id { get; set; }

        [Required]
        public string Login { get; set; }
		[Required]
		public string Password { get; set; }
		[Required]
		public string Role { get; set; }

        public RefreshToken RefreshToken { get; set; }
        public Profile Profile { get; set; }
        public IList<Event> OrganizedEvents { get; set; } = new List<Event>();
        public IList<Event> VisitedEvents { get; set; } = new List<Event>();
    }
}
