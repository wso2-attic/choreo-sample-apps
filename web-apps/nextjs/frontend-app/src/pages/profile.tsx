import Layout from "@/components/Layouts";
import { useSession } from "next-auth/react";

function ProfilePage() {
  const { data: session } = useSession();
  if (!session) {
    return null;
  }
  return (
    <Layout>
      <h1>This is the Profile Page</h1>
      <p>Signed in as {session?.user?.email}</p>
      {session.expires && <p>Expires: {session.expires}</p>}
    </Layout>
  );
}

export default ProfilePage;
