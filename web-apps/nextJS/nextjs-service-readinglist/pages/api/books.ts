import { NextApiRequest, NextApiResponse } from "next";
import prisma from "../../lib/prisma";
import jwt_decode from "jwt-decode";

interface DecodedToken {
  email: any;
}

export default async function handle(
  req: NextApiRequest,
  res: NextApiResponse
) {
  try {
    if (req.method === "POST") {
      const accessToken = req.headers.authorization;

      if (typeof accessToken !== "string") {
        throw new Error("Access token is missing or invalid");
      }

      const decodedToken = jwt_decode(accessToken) as DecodedToken;

      const requestBody = JSON.parse(req.body);
      console.log(requestBody.title);
      const newBook = await prisma.book.create({
        data: {
          title: requestBody.title,
          author: requestBody.author,
          status: requestBody.status,
          user_writer_id: decodedToken.email,
        },
      });
      res.setHeader("Access-Control-Allow-Origin", "*");
      res.setHeader(
        "Access-Control-Allow-Methods",
        "GET, POST, PUT, DELETE, OPTIONS"
      );
      res.setHeader("Access-Control-Allow-Headers", "*");
      res.send(JSON.stringify(newBook));
    } else if (req.method === "DELETE") {
      const bookId = req.query.id;

      if (typeof bookId !== "string") {
        throw new Error("Invalid book ID");
      }

      const deletedBook = await prisma.book.delete({
        where: { id: bookId },
      });
      res.setHeader("Access-Control-Allow-Origin", "*");
      res.setHeader(
        "Access-Control-Allow-Methods",
        "GET, POST, PUT, DELETE, OPTIONS"
      );
      res.setHeader("Access-Control-Allow-Headers", "*");
      res.json(deletedBook);
    } else if (req.method === "GET") {
      const accessToken = req.headers.authorization;

      if (typeof accessToken !== "string") {
        throw new Error("Access token is missing or invalid");
      }

      const decodedToken = jwt_decode(accessToken) as DecodedToken;

      try {
        const reads = await prisma.book.findMany({
          where: {
            user_writer_id: decodedToken.email,
          },
        });
        res.setHeader("Access-Control-Allow-Origin", "*");
        res.setHeader(
          "Access-Control-Allow-Methods",
          "GET, POST, PUT, DELETE, OPTIONS"
        );
        res.setHeader("Access-Control-Allow-Headers", "*");
        res.send(JSON.stringify(reads));
      } catch (error) {
        res.setHeader("Access-Control-Allow-Origin", "*");
        res.setHeader(
          "Access-Control-Allow-Methods",
          "GET, POST, PUT, DELETE, OPTIONS"
        );
        res.setHeader("Access-Control-Allow-Headers", "*");
        res.status(500).end();
      }
    }
    if (req.method === "OPTIONS") {
      res.setHeader("Access-Control-Allow-Origin", "*");
      res.setHeader(
        "Access-Control-Allow-Methods",
        "GET, POST, PUT, DELETE, OPTIONS"
      );
      res.setHeader("Access-Control-Allow-Headers", "*");
      res.status(200).end();
      return;
    }
  } catch (error) {
    console.error("Error occurred", error);
    res.setHeader("Access-Control-Allow-Origin", "*");
    res.setHeader(
      "Access-Control-Allow-Methods",
      "GET, POST, PUT, DELETE, OPTIONS"
    );
    res.setHeader("Access-Control-Allow-Headers", "*");
    res.status(500).end();
  }
}
