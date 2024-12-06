using api.Entities;
using Microsoft.EntityFrameworkCore;

namespace api.Database
{
    public class DatabaseContext : DbContext
    {
        public DbSet<User> Users { get; set; } = null!;
        public DbSet<Profile> Profiles { get; set; } = null!;
        public DbSet<RefreshToken> RefreshTokens { get; set; } = null!;
        public DbSet<Event> Events { get; set; } = null!;
        public DbSet<EventAdditional> EventAdditionals { get; set; } = null!;
        public DbSet<EventGeneral> EventGenerals { get; set; } = null!;

        public DatabaseContext(DbContextOptions<DatabaseContext> options)
            : base(options)
        {

        }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<User>()
                .HasMany(user => user.OrganizedEvents)
                .WithOne(e => e.Organizer)
                .HasForeignKey(e => e.OrganizerId)
                .OnDelete(DeleteBehavior.NoAction);
            modelBuilder.Entity<User>()
                .HasMany(user => user.VisitedEvents)
                .WithMany(e => e.Visiters);
        }
    }
}
