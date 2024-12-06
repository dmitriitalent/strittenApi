namespace api.Exceptions;

public class UnauthorizedException<TController> : ResponseException
{
	public UnauthorizedException(string message = "Пользователь не авторизован")
		: base(message, 401, typeof(TController).Name)
	{ }
}
