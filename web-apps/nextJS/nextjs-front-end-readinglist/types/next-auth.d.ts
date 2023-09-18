import { JWT } from "next-auth/jwt";
import NextAuth, { DefaultSession } from "next-auth";

declare module "next-auth/jwt" {
  interface JWT {
    idToken?: string;
    accessToken?: string;
  }
}

declare module "next-auth" {
  interface Session {
    user: {
      idToken?: string;
      accessToken?: string;
    } & DefaultSession["user"];
  }

  interface User {
    idToken?: string;
    accessToken?: string;
  }
}