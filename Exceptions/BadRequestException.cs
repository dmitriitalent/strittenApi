namespace api.Exceptions;

public class BadRequestException<TController> : ResponseException
{
	public BadRequestException(string message)
		: base(message, 400, typeof(TController).Name) 
	{ }
}
