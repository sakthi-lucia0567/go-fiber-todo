import { Box, List, Notification, ThemeIcon } from "@mantine/core";
import useSWR from "swr";
import AddTodo from "./components/AddTodo";
import { CheckCircleFillIcon } from "@primer/octicons-react";
import { useState } from "react";

export const ENDPOINT = "https://go-fiber-todo.onrender.com";

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

export interface Todo {
  id: number;
  title: string;
  body: string;
  done: boolean;
}

function App() {
  const { data, mutate } = useSWR<Todo[]>("api/todos", fetcher);

  const [success, isSuccess] = useState(false);

  async function markTodoAsDone(id: number) {
    const updated = await fetch(`${ENDPOINT}/api/todos/${id}/done`, {
      method: "PATCH",
    }).then((r) => r.json);
    isSuccess(true);
    setTimeout(closeNotification, 5000);
    mutate(updated);
  }

  function closeNotification() {
    isSuccess(false);
  }

  return (
    <>
      <Box p="2rem" w="100%" maw={"40rem"} m={"0 auto"}>
        <List spacing="xs" size="sm" mb={12} center>
          {data?.map((todo) => {
            return (
              <>
                <List.Item
                  key={`todo_list_${todo.id}`}
                  onClick={() => markTodoAsDone(todo.id)}
                  icon={
                    todo.done ? (
                      <ThemeIcon color="teal" size={24} radius={"xl"}>
                        <CheckCircleFillIcon fill="green" size={20} />
                      </ThemeIcon>
                    ) : (
                      <ThemeIcon color="gray" size={24} radius={"xl"}>
                        <CheckCircleFillIcon size={20} />
                      </ThemeIcon>
                    )
                  }
                >
                  {todo.title}
                </List.Item>
              </>
            );
          })}
        </List>
        <AddTodo mutate={mutate} />
      </Box>
      {success && (
        <Notification
          w={"20%"}
          withBorder
          color="gray"
          title="Todo Updated"
          onClose={() => closeNotification()}
          closeButtonProps={{ "aria-label": "Hide notification" }}
        />
      )}
    </>
  );
}

export default App;
