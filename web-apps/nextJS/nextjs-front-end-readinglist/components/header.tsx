import { getSession, signIn, signOut } from "next-auth/react";

const Header = (session: any, org: any) => (
  <header>
    <div className={`max-w-4xl mx-auto py-16 px-14 sm:px-6 lg:px-8`}>
      <h1
        className={`font-sans font-bold text-4xl md:text-5xl lg:text-8xl text-center leading-snug text-gray-800`}
      >
        Reading List Application
      </h1>
      <div className={`max-w-xl mx-auto`}>
        <p className={`mt-10 text-gray-900 text-center text-xl lg:text-3xl`}>
          Add books to your reading list
        </p>
      </div>
      <div className={`mt-10 flex justify-center items-center w-full mx-auto`}>
        {session ? (
          <button
            className="bg-white bg-opacity-20 p-2 rounded-md text-sm my-3 font-medium text-white h-10"
            onClick={(e) => {
              e.preventDefault();
              signIn("asgardeo", { callbackUrl: "/" });
            }}
          >
            Login with Asgardeo
          </button>
        ) : (
          <button
            onClick={(e) => {
              e.preventDefault();
              signOut({ callbackUrl: "/" });
            }}
            style={{ marginTop: "30px", marginBottom: "40px" }}
          >
            Logout
          </button>
        )}
      </div>
    </div>
  </header>
);

export default Header;

export async function getServerSideProps(context: any) {
  return {
    props: {
      session: await getSession(context),
      org: process.env.ASGARDEO_ORGANIZATION_NAME,
    },
  };
}
