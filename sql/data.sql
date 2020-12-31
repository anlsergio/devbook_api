INSERT INTO users (
  name,
  username,
  email,
  password
) VALUES
(
  "Peter Griffin",
  "pgriffin",
  "pgriffin@gmail.com",
  "$2a$10$TtXrhGatQ2y4z3hMcUpjm.jozgnQDjyaFQoYl045sdp3BAsBCQnAe"
),
(
  "Brian Griffin",
  "bark4girls",
  "bark4girls@gmail.com",
  "$2a$10$TtXrhGatQ2y4z3hMcUpjm.jozgnQDjyaFQoYl045sdp3BAsBCQnAe"
),
(
  "Joe Swanson",
  "jswanson",
  "jswanson@yahoo.com",
  "$2a$10$TtXrhGatQ2y4z3hMcUpjm.jozgnQDjyaFQoYl045sdp3BAsBCQnAe"
);

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(1, 3),
(2, 3),
(3, 2);

INSERT INTO posts (
  title,
  content,
  author_id
) VALUES (
  "Peter's Post",
  "This is a random post from Peter!",
  1
),
(
  "Brian's Post",
  "This is a random post from Brian!",
  2
),
(
  "Joe's Post",
  "This is a random post from Joe!",
  3
);