using api.Entities;
using System.ComponentModel.DataAnnotations;

namespace api.Dtos
{
    public class GetUserDTO
    {
        public int Id { get; set; }
        public string Login { get; set; }
        public string Role { get; set; }

        public GetUserDTOProfile Profile { get; set; }
    }
    public class GetUserDTOProfile
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string Surname { get; set; }
        public string Email { get; set; }
        public int UserId { get; set; }
    }
}

