import Layout from "@/components/Layouts";
import { listTodosForUser, TodoResponse } from "@/svc/backend.client";
import { GetServerSideProps, InferGetServerSidePropsType } from "next";
import { getServerSession } from "next-auth";
import { getSession, SessionContext } from "next-auth/react";
import Link from "next/link";
import { authOptions } from "@/pages/api/auth/[...nextauth]";
import { redirectTo, redirectToLogin } from "@/utils/redirect";
import { getNextAuthServerSession } from "@/utils/session";

type TodosPageProps = {
  todos: TodoResponse[];
};

export const getServerSideProps = (async (context: any) => {
  const session = await getNextAuthServerSession(context);
  if (!session) {
    return redirectToLogin();
  }
  const todos = await listTodosForUser((session as any)?.user?.id!);
  return { props: { todos } };
}) satisfies GetServerSideProps<TodosPageProps>;

function TodosPage({
  todos,
}: InferGetServerSidePropsType<typeof getServerSideProps>) {
  return (
    <Layout>
      <div>
        <h1>Todos List</h1>
        <ul>
          {todos.map((todo) => (
            <li key={todo.id}>
              <Link href={`/todo/${todo.id}`}>{todo.title}</Link>
            </li>
          ))}
        </ul>
      </div>
    </Layout>
  );
}

export default TodosPage;
