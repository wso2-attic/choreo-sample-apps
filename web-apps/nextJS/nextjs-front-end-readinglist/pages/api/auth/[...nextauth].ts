import NextAuth from "next-auth"
import getConfig from "next/config"

const { serverRuntimeConfig, publicRuntimeConfig } = getConfig()
export default NextAuth({

  providers: [

    {
      id: "asgardeo",
      name: "Asgardeo",
      clientId: serverRuntimeConfig.ASGARDEO_CLIENT_ID,
      clientSecret: serverRuntimeConfig.ASGARDEO_CLIENT_SECRET,
      type: "oauth",
      wellKnown: "https://api.asgardeo.io/t/" + publicRuntimeConfig.ASGARDEO_ORGANIZATION_NAME + "/oauth2/token/.well-known/openid-configuration",
      authorization: { params: { scope: publicRuntimeConfig.ASGARDEO_SCOPES } },
      idToken: true,
      checks: ["pkce", "state"],
      profile(profile) {
        return {
          id: profile.sub,
          name: profile.name,
          email: profile.email,
        }
      },
    },
  ],
  secret: serverRuntimeConfig.SECRET,

  session: {
    strategy: "jwt",
  },
  callbacks: {
    async session({ session, token, user }) {
      session.user.accessToken = token.accessToken
      session.user.idToken = token.idToken
      return session
    },
    async jwt({ token, user, account, profile, isNewUser }) {
      if (account) {
        token.accessToken = account.access_token
        token.idToken = account.id_token
      }
      return token
    }
  },

  theme: {
    colorScheme: "light",
  },

  debug: true,
})