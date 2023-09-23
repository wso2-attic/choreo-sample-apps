import NextLink from "next/link";
import { Link } from "@nextui-org/link";
import { Snippet } from "@nextui-org/snippet";
import { Code } from "@nextui-org/code";
import { button as buttonStyles } from "@nextui-org/theme";
import { siteConfig } from "@/config/site";
import { title, subtitle } from "@/components/primitives";
import DefaultLayout from "@/layouts/default";
import { getServerSession } from "next-auth";
import { useSession, signIn, signOut, getSession } from "next-auth/react";
import authOptions from "@/pages/api/auth/[...nextauth]";
import { GetServerSideProps } from "next";
import { redirect } from "next/dist/server/api-utils";
import { redirectTo, redirectToLogin } from "@/utils/redirect";
import { getNextAuthServerSession } from "@/utils/session";
import { Button } from "@nextui-org/react";
import { FcAddRow } from "react-icons/fc";

export const getServerSideProps = (async (context: any) => {
  const session = await getNextAuthServerSession(context);
  if (!session) {
    return redirectToLogin();
  }
  return { props: {} };
}) satisfies GetServerSideProps<{}>;

export default function IndexPage() {
  return (
    <DefaultLayout>
      <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
        <div className="inline-block max-w-lg text-center justify-center">
          <h1 className={title({ color: "cyan" })}>Make your&nbsp;</h1>
          <h1 className={title({ color: "blue" })}>ToDo Lists&nbsp;</h1>
          <br />
          <h1 className={title({ color: "cyan" })}>here.</h1>
          <h4 className={subtitle({ class: "mt-4" })}>
            Add, organize and view your Todos.
          </h4>
        </div>
      </section>
    </DefaultLayout>
  );
}
