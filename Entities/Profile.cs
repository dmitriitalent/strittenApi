using System.ComponentModel.DataAnnotations;

namespace api.Entities
{
    public class Profile
    {
        public int Id { get; set; }
        [Required]
        public string Name { get; set; }
        [Required]
        public string Surname { get; set; }
        [Required]
        public string Email { get; set; }

        public int UserId { get; set; }
        public User User { get; set; }
    }
}
