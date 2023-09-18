import { useRouter } from "next/router";
import { getSession, signOut, useSession } from 'next-auth/react'

export default function SignOutButton() {
  const router = useRouter();
  const { data: session, status: sessionStatus } = useSession();

  async function signOutHandler() {

    await fetch(`/api/auth/federated-sign-out`, {
      method: "POST",
      body: JSON.stringify({
        idToken: session?.user.idToken,
      }),
    });
/*
    router.push({
      pathname: '/api/auth/federated-sign-out',
      query: { idToken: session?.user.idToken }
  })
  */
  }

  return (
    <button type="button"  onClick={() => {
      signOutHandler();
      signOut({ redirect: false }).then(() => {
        router.push("/");
      });
    }}>
      Sign out
    </button>
  );
}

export async function getServerSideProps(context: any) {
  return {
    props: {
      session: await getSession(context),
    },
  };
}