namespace api.Exceptions;

public class NotFoundException<TController> : ResponseException
{
	public NotFoundException(string message = "NotFound")
		: base(message, 404, typeof(TController).Name)
	{ }
}