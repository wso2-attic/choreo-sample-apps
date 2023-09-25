import Link from "next/link";
import styles from "./Layout.module.css";
import { signOut, useSession } from "next-auth/react";
import { redirectTo } from "@/utils/redirect";

export default function Layout({ children }: { children: React.ReactNode }) {
  const { status } = useSession({ required: true });
  return (
    <div className={styles.container}>
      <nav className={styles.sideMenu}>
        <ul>
          <li>
            <Link href="/">Home</Link>
          </li>
          <li>
            <Link href="/profile">Profile</Link>
          </li>
          <li>
            <Link href="/todos">Todos</Link>
          </li>
          <br></br>
          <br></br>
          <br></br>
          <br></br>
          <br></br>
          <User />
        </ul>
      </nav>
      <div className={styles.content}>{children}</div>
    </div>
  );
}

function User() {
  const { data: session } = useSession();
  if (!session) {
    return null;
  }
  return (
    <>
      Signed in as {session?.user?.email} <br />
      <button onClick={() => signOut()}>Sign out</button>
    </>
  );
}
