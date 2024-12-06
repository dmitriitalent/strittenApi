using System.ComponentModel.DataAnnotations;

namespace api.Entities
{
    public class RefreshToken
    {
        public int Id { get; set; }
        [Required]
        public string Token { get; set; }
		[Required]
		public int UserId { get; set; }
        [Required]
        public User User { get; set; }
    }
}
