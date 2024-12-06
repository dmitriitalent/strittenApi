using Microsoft.Identity.Client;
using Microsoft.IdentityModel.Tokens;
using System.Diagnostics;
using System.IdentityModel.Tokens.Jwt;
using System.Runtime.CompilerServices;
using System.Security.Claims;
using System.Text;
using api;

namespace adasdasd.Services;

public class Legacy
{
	public string GenerateRefreshToken(
		IList<Claim> payload
	)
	{
		JwtSecurityToken RefreshJWT = new JwtSecurityToken(
			claims: payload,
			audience: TokenSettings.Audience,
			issuer: TokenSettings.Issuer,
			expires: DateTime.Now.AddDays(TokenSettings.RefreshTokenLifeTimeInDays),
			signingCredentials: new SigningCredentials(TokenSettings.GetSymmetricSecurityRefreshTokenKey(), SecurityAlgorithms.HmacSha256)
		);

		return new JwtSecurityTokenHandler().WriteToken(RefreshJWT);
	}
	public string GenerateAccessToken(
		IList<Claim> payload
	)
	{
		JwtSecurityToken AccessJWT = new JwtSecurityToken(
			claims: payload,
			audience: TokenSettings.Audience,
			issuer: TokenSettings.Issuer,
			expires: DateTime.Now.AddMinutes(TokenSettings.AccessTokenLifeTimeInMinutes),
			signingCredentials: new SigningCredentials(TokenSettings.GetSymmetricSecurityAccessTokenKey(), SecurityAlgorithms.HmacSha256)
		);

		return new JwtSecurityTokenHandler().WriteToken(AccessJWT);
	}

	public bool TryVerifyRefreshToken(string refreshToken)
	{
		JwtSecurityToken RefreshJWT = new JwtSecurityToken(refreshToken);
		new JwtSecurityTokenHandler().ValidateToken(refreshToken, TokenSettings.validationParameters, out SecurityToken Token);

		return true;
	}

	public int GetId(string RefreshJWT)
	{
		return int.Parse(new JwtSecurityTokenHandler().ReadJwtToken(RefreshJWT).Claims.FirstOrDefault(claim => claim.Type == "Id").Value);
	}
}



public static class TokenSettings
{
	public static string Issuer = "Server";
	public static string Audience = "Client";
	public static int RefreshTokenLifeTimeInDays = 30;
	public static int AccessTokenLifeTimeInMinutes = 30;
	public static string RefreshTokenSecurityKey = "RefreshSecurityKey";
	public static string AccessTokenSecurityKey = "AccessSecurityKey";

	public static TokenValidationParameters validationParameters = new TokenValidationParameters
	{
		ValidateIssuer = true,
		ValidateAudience = true,
		ValidateLifetime = true,
		ValidateIssuerSigningKey = true,
		ValidIssuer = Issuer,
		ValidAudience = Audience,
		IssuerSigningKey = GetSymmetricSecurityRefreshTokenKey()
	};

	public static SymmetricSecurityKey GetSymmetricSecurityRefreshTokenKey() =>
		new SymmetricSecurityKey(Encoding.UTF8.GetBytes(RefreshTokenSecurityKey));
	public static SymmetricSecurityKey GetSymmetricSecurityAccessTokenKey() =>
		new SymmetricSecurityKey(Encoding.UTF8.GetBytes(AccessTokenSecurityKey));
}
