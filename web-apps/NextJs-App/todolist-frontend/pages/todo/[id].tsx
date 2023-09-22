import Layout from "@/components/Layouts";
import { useRouter } from "next/router";

function TodoDetailPage() {
  const router = useRouter();
  const { id } = router.query;

  // You'd fetch data here based on the ID or handle it differently.
  // For demonstration purposes, let's just display the ID.

  return (
    <Layout>
      <h1>Todo Detail for ID: {id}</h1>
    </Layout>
  );
}

export default TodoDetailPage;
