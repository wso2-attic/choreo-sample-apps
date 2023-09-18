import { useSession, signOut, getSession } from "next-auth/react";
import { Tab } from "@headlessui/react";
import { useEffect, useState } from "react";
import Router from "next/router";
import { Dictionary } from "lodash";
import groupBy from "lodash/groupBy";
import { Book } from "@/types/book";
import AddItem from "@/components/add-book";
import Header from "@/components/header";
import getConfig from "next/config";

// import { getReadingListInstance } from "@/lib/instance";

export function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(" ");
}

const { serverRuntimeConfig, publicRuntimeConfig } = getConfig()

const Home = () => {
  const { data: session, status } = useSession();
  const [readingList, setReadingList] = useState<Dictionary<Book[]> | null>(
    null
  );
  const [isAddItemOpen, setIsAddItemOpen] = useState(false);

  if (status === "loading") {
    return <div>Authenticating...</div>;
  }

  const accessToken = session?.user?.accessToken;

  async function getBooks() {
try{
  const response = await fetch(`${publicRuntimeConfig.NEXT_PUBLIC_SERVICE_URL}`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${accessToken}`,
    },
  });

  if (response.ok) {
    const grouped = groupBy(await response.json(), "status");
    setReadingList(grouped);
    Router.push("/");
  } else {
    alert(`Something went wrong!"`);
  }
} catch (e) {
  signOut({ callbackUrl: "/" });
}
  
  }
  async function deleteBook(id: string): Promise<void> {
    await fetch(`${process.env.NEXT_PUBLIC_SERVICE_URL}?id=${id}`, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });
    alert("Book is deleted successfully");
    Router.push("/");
  }

  useEffect(() => {
    if (accessToken) {
      try {
        getBooks();
      } catch (e) {
        signOut({ callbackUrl: "/" });
      }
    }
  }, [accessToken]);

  return (
    <div className="py-3 md:py-6">
      <div className="container px-4 mx-auto flex justify-left">
        <div className="w-full px-1 py-16 sm:px-0 mb-20">
          {session?.user ? (
            <>
              <nav>
                <div className="flex justify-between">
                  <p className="text-4xl text-white mb-3 font-bold">
                    Reading List
                  </p>
                  <p className="text-3xl text-white mb-3">
                    Hello {session?.user.email ? session.user.email : ""}
                  </p>
                  <button
                    className="bg-black bg-opacity-20 p-2 rounded-md text-sm mb-3 font-medium text-white h-10"
                    onClick={(e) => {
                      e.preventDefault();
                      signOut({ callbackUrl: "/" });
                    }}
                  >
                    Logout
                  </button>
                </div>
                <div className="flex justify-between">
                  <button
                    className="bg-black bg-opacity-20 p-2 rounded-md text-sm my-3 font-medium text-white h-10"
                    onClick={() => setIsAddItemOpen(true)}
                  >
                    + Add New
                  </button>
                </div>
              </nav>
              {readingList && (
                <Tab.Group>
                  <Tab.List className="flex space-x-1 rounded-xl bg-blue-900/20 p-1">
                    {Object.keys(readingList).map((val) => (
                      <Tab
                        key={val}
                        className={({ selected }) =>
                          classNames(
                            "w-full rounded-lg py-2.5 text-sm font-medium leading-5 text-blue-700",
                            "ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2",
                            selected
                              ? "bg-white shadow"
                              : "text-blue-100 hover:bg-white/[0.12] hover:text-white"
                          )
                        }
                      >
                        {val}
                      </Tab>
                    ))}
                  </Tab.List>
                  <Tab.Panels className="mt-2">
                    {Object.values(readingList).map((books: Book[], idx) => (
                      <Tab.Panel key={idx}>
                        <ul>
                          {books.map((book) => (
                            <div className="flex justify-between" key={book.id}>
                              <li className="relative rounded-md p-3">
                                <h3 className="text-sm font-medium leading-5">
                                  {book.title}
                                </h3>
                                <h6 className="text-xs font-normal leading-4 text-gray-900">
                                  <i> Written By: </i>
                                  {book.author}
                                </h6>
                              </li>
                              <button
                                className="float-right bg-red-500 text-white rounded-md self-center text-xs p-2 mr-2"
                                onClick={() => deleteBook(book.id || "")}
                              >
                                Delete
                              </button>
                            </div>
                          ))}
                        </ul>
                      </Tab.Panel>
                    ))}
                  </Tab.Panels>
                </Tab.Group>
              )}
              <AddItem isOpen={isAddItemOpen} setIsOpen={setIsAddItemOpen} />
            </>
          ) : (
            <>
              <Header />
            </>
          )}
        </div>
      </div>
    </div>
  );
};

export async function getServerSideProps(context: any) {
  return {
    props: {
      session: await getSession(context),
    },
  };
}

export default Home;
