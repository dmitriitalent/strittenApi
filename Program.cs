using Microsoft.EntityFrameworkCore;
using api.Database;
using api.Exceptions;

var builder = WebApplication.CreateBuilder(args);


string connection = builder.Configuration.GetConnectionString("DefaultConnection");

builder.Services.AddDbContext<DatabaseContext>(options => options.UseSqlServer(connection));
builder.Services.AddCors(options =>
{
    options.AddPolicy(name: "AllCors",
        policy =>
        {
            policy.WithOrigins("http://localhost:3000")
                  .AllowAnyMethod()
                  .AllowAnyHeader()
                  .AllowCredentials();
        });
});
builder.Services.AddControllers();

var app = builder.Build();

app.UseCors("AllCors");

app.UseHttpsRedirection();

app.UseAuthorization();

app.Use(async (context, next) =>
{
    try
    {
        await next.Invoke();
    }
    catch (ResponseException ex)
    {
        context.Response.StatusCode = ex.StatusCode;
        await context.Response.WriteAsync(ex.Message);
    }
});

app.MapControllers();

app.Run();
