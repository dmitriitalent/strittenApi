namespace api.Exceptions;

public class ResponseException : Exception
{
	public ResponseException(string message, int statusCode, string controllerName) 
		: base(message) 
	{
		this.StatusCode = statusCode;
		this.Controller = controllerName;
	}

	public int StatusCode { get; set; }
	public string? Controller { get; set; }
}

