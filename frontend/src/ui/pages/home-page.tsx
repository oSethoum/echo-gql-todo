import {
  Box,
  Button,
  Group,
  Paper,
  ScrollArea,
  Switch,
  Text,
  TextInput,
} from "@mantine/core";
import { useForm, FormErrors } from "@mantine/form";
import { useEffect, useState } from "react";
import {
  Event,
  useCreateTodoMutation,
  useTodosQuery,
  useUpdateTodoMutation,
  useTodosSubSubscription,
  useDeleteTodoMutation,
} from "../../graphql";

export function HomePage() {
  const form = useForm({
    initialValues: {
      text: "",
    },
  });

  const [todosResponse] = useTodosQuery();
  const [createdTodo, createTodo] = useCreateTodoMutation();
  const [updatedTodo, updateTodo] = useUpdateTodoMutation();
  const [deletedTodo, deleteTodo] = useDeleteTodoMutation();
  const [create] = useTodosSubSubscription({
    variables: {
      event: Event.Create,
    },
  });

  const [totalCount, setTotalCount] = useState(
    todosResponse.data?.todos.totalCount || 0
  );

  const [update] = useTodosSubSubscription({
    variables: {
      event: Event.Update,
      query: {
        where: {
          done: true,
        },
      },
    },
  });

  const [totalDone, setTotalDone] = useState(
    todosResponse.data?.todos.totalCount || 0
  );

  useEffect(() => {
    console.log("Create Event ", create);
    setTotalCount((count) => create.data?.todos.totalCount || count);
  }, [create]);

  useEffect(() => {
    console.log("Update Event ", update);
    setTotalDone((count) =>
      update.data?.todos.totalCount != undefined
        ? update.data?.todos.totalCount
        : count
    );
  }, [update]);

  return (
    <Box
      m="lg"
      sx={{
        display: "flex",
        justifyContent: "center",
        flexDirection: "column",
      }}
    >
      <Box sx={{ width: 400 }}>
        <form
          onSubmit={form.onSubmit((values) => {
            createTodo({
              input: {
                text: values.text,
              },
            });
            form.reset();
          })}
        >
          <Box sx={{ display: "flex", gap: 5 }}>
            <TextInput
              sx={{ flexGrow: 1 }}
              placeholder="new todo"
              {...form.getInputProps("text")}
            />
            <Button type="submit">add</Button>
          </Box>
        </form>
        <Group my={5}>
          <Text>Count: {totalCount}</Text>
          <Text>Done: {totalDone}</Text>
        </Group>
        <ScrollArea offsetScrollbars sx={{ height: 600 }}>
          {todosResponse.data?.todos.edges?.map((e, index) => (
            <Paper
              withBorder
              p={5}
              sx={{
                display: "flex",
                gap: 10,
              }}
              key={index}
            >
              <Box>{e?.node?.id}</Box>
              <Box sx={{ flexGrow: 1 }}>{e?.node?.text}</Box>
              <Switch
                checked={e?.node?.done}
                onChange={(event) => {
                  updateTodo({
                    id: e?.node?.id as string,
                    input: {
                      done: event.target.checked,
                    },
                  });
                }}
              />
              <Button
                color="red"
                onClick={() => {
                  deleteTodo({ id: e?.node?.id as string });
                }}
              >
                Delete
              </Button>
            </Paper>
          ))}
        </ScrollArea>
      </Box>
    </Box>
  );
}

export default HomePage;
