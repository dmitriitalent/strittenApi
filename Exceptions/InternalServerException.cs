using Microsoft.AspNetCore.Mvc;

namespace api.Exceptions;

public class InternalServerException<TController> : ResponseException
{
	public InternalServerException(string message)
		: base(message, 500, typeof(TController).Name)
	{ }
}
