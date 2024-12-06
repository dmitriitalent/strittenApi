using api.Entities;

namespace api.Dtos
{
    public class EventVisitersDTO : List<UserDTO>
    {
        public EventVisitersDTO(IList<User> EventVisiters, EventDTO parentEvent = null) 
        {
            if (EventVisiters != null)
            {
                foreach (var User in EventVisiters)
                {
                    this.Add(new UserDTO(User, parentEvent));
                }
            }
        }
    }
}
