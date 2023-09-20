import Layout from "@/components/Layouts";
import { getServerSession } from "next-auth";
import { useSession, signIn, signOut, getSession } from "next-auth/react";
import Link from "next/link";
import authOptions from "@/pages/api/auth/[...nextauth]";
import { GetServerSideProps } from "next";
import { redirect } from "next/dist/server/api-utils";
import { redirectTo, redirectToLogin } from "@/utils/redirect";
import { getNextAuthServerSession } from "@/utils/session";

export const getServerSideProps = (async (context: any) => {
  const session = await getNextAuthServerSession(context);
  if (!session) {
    return redirectToLogin();
  }
  return { props: {} };
}) satisfies GetServerSideProps<{}>;

export default function HomePage() {
  return (
    <Layout>
      <div>
        <h1>Welcome to the Home Page</h1>
        <ul>
          <li>
            <Link href="/profile">Profile</Link>
          </li>
          <li>
            <Link href="/todos">Todos</Link>
          </li>
        </ul>
      </div>
    </Layout>
  );
}
