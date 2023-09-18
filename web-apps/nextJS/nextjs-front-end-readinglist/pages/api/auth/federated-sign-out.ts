import { NextApiRequest, NextApiResponse } from "next";
import { getSession } from "next-auth/react";
import { useRouter } from "next/router";

export default async function federatedSignOut(
  req: NextApiRequest,
  res: NextApiResponse
) {
  // Get the site base url.
  const baseUrl = "https://fd9524ea-680d-45af-9efa-7cbd6329585c.e1-us-east-azure.choreoapps.dev";

  const router = useRouter();

  try {
    const requestBody = JSON.parse(req.body);
    const idToken = requestBody.idToken;
    const session = await getSession({ req });
    if (!session) {
      return res.redirect(baseUrl);
    }

    // Asgardeo logout endpoint.
    const endSessionURL = `https://api.asgardeo.io/t/areeb/oidc/logout`;

    const redirectURL = `https://fd9524ea-680d-45af-9efa-7cbd6329585c.e1-us-east-azure.choreoapps.dev/`;

    const endSessionParams = new URLSearchParams({
      id_token_hint: idToken,
      post_logout_redirect_uri: redirectURL,
    });
    const fullUrl = `${endSessionURL}?${endSessionParams.toString()}`;
    return res.redirect(fullUrl);
  } catch (error) {
    res.redirect(baseUrl);
  }
}
