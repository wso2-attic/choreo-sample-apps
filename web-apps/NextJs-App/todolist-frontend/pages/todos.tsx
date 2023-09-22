// import Layout from "@/components/Layouts";
import { listTodosForUser, TodoResponse } from "@/svc/backend.client";
import { GetServerSideProps, InferGetServerSidePropsType } from "next";
import { getServerSession } from "next-auth";
import { getSession, SessionContext } from "next-auth/react";
import Link from "next/link";
import { authOptions } from "@/pages/api/auth/[...nextauth]";
import { redirectTo, redirectToLogin } from "@/utils/redirect";
import { getNextAuthServerSession } from "@/utils/session";
import DefaultLayout from "@/layouts/default";
import {
  Chip,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from "@nextui-org/react";

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
    <DefaultLayout>
      <Table aria-label="Example table with dynamic content">
        <TableHeader>
          <TableColumn key="col">TODO List</TableColumn>
          <TableColumn key="col">Status</TableColumn>
        </TableHeader>
        <TableBody emptyContent={"No rows to display."}>
          {todos.map((todo) => (
            <TableRow key={todo.id}>
              <TableCell>
                <Link href={`/todo/${todo.id}`}>{todo.title}</Link>
              </TableCell>
              <TableCell>
                <Chip
                  className="capitalize"
                  color="success"
                  size="sm"
                  variant="flat"
                >
                  Active
                </Chip>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </DefaultLayout>
  );
}

export default TodosPage;
