import {
  createTodoForUser,
  deleteTodoForUser,
  listTodosForUser,
  TodoResponse,
} from "@/svc/backend.client";
import { GetServerSideProps, InferGetServerSidePropsType } from "next";
import Link from "next/link";
import { redirectTo, redirectToLogin } from "@/utils/redirect";
import { getNextAuthServerSession } from "@/utils/session";
import DefaultLayout from "@/layouts/default";
import {
  Button,
  Chip,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  Input,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
  useDisclosure,
} from "@nextui-org/react";
import { useState } from "react";
import { FcPlus } from "react-icons/fc";
import { MdDelete } from "react-icons/md";

type TodosPageProps = {
  todos: TodoResponse[];
  userId: string;
};

export const getServerSideProps = (async (context: any) => {
  const session = await getNextAuthServerSession(context);
  if (!session) {
    return redirectToLogin();
  }
  const todos = await listTodosForUser((session as any)?.user?.id!);
  const userId = (session as any)?.user?.id || "";
  return { props: { todos, userId } };
}) satisfies GetServerSideProps<TodosPageProps>;

function TodosPage({
  todos,
  userId,
}: InferGetServerSidePropsType<typeof getServerSideProps>) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");

  const onSave = async () => {
    if (!userId) {
      redirectToLogin();
    } else {
      try {
        const todo = await createTodoForUser(userId as string, {
          title,
          description,
        });
        console.log("Todo created:", todo);
      } catch (e) {
        alert("Something went wrong!");
      } finally {
        onOpenChange();
      }
    }
  };

  const onDelete = async (todoId: number) => {
    if (!userId) {
      redirectToLogin();
    } else {
      try {
        const todo = await deleteTodoForUser(userId as string, todoId);
        console.log("Todo deleted:");
      } catch (e) {
        alert("Something went wrong!");
      }
    }
  };
  return (
    <DefaultLayout>
      <div className="inline-block max-w-lg text-center justify-center my-4">
        <Button onPress={onOpen} color="secondary" variant="ghost">
          <FcPlus />
          Add Todo
        </Button>
      </div>
      <Modal
        backdrop="opaque"
        isOpen={isOpen}
        onOpenChange={onOpenChange}
        radius="lg"
        classNames={{
          body: "py-6",
          backdrop: "bg-[#292f46]/50 backdrop-opacity-40",
          base: "border-[#292f46] bg-[#19172c] dark:bg-[#19172c] text-[#a8b0d3]",
          header: "border-b-[1px] border-[#292f46]",
          footer: "border-t-[1px] border-[#292f46]",
          closeButton: "hover:bg-white/5 active:bg-white/10",
        }}
      >
        <ModalContent>
          {(onClose) => (
            <>
              <ModalHeader className="flex flex-col gap-1">
                Add Todo
              </ModalHeader>
              <ModalBody>
                <Input
                  type="text"
                  label="Title"
                  onChange={(e) => setTitle(e.target.value)}
                />

                <Input
                  type="text"
                  label="Description"
                  onChange={(e) => setDescription(e.target.value)}
                />
              </ModalBody>
              <ModalFooter>
                <Button color="primary" variant="light" onPress={onClose}>
                  Close
                </Button>
                <Button
                  className="bg-[#6f4ef2] shadow-lg shadow-indigo-500/20"
                  onPress={onSave}
                >
                  Action
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
      <Table aria-label="Example table with dynamic content">
        <TableHeader>
          <TableColumn key="col">TODO List</TableColumn>
          <TableColumn key="col">Status</TableColumn>
          <TableColumn key="col">Delete</TableColumn>
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
              <TableCell>
                <Button
                  onPress={() => {
                    onDelete(todo.id);
                  }}
                  variant="light"
                >
                  <MdDelete color="red" size={20} />
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </DefaultLayout>
  );
}

export default TodosPage;
