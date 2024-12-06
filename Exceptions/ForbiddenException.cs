namespace api.Exceptions;

public class ForbiddenException<TController> : ResponseException
{
	public ForbiddenException(string message = "Forbidden")
		: base(message, 403, typeof(TController).Name)
	{ }
}