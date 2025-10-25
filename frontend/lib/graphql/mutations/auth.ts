import { gql } from '@apollo/client';

export const REGISTER_USER = gql`
  mutation RegisterUser($input: RegisterUserInput!) {
    registerUser(input: $input) {
      id
      name
      email
    }
  }
`;

export const LOGIN_USER = gql`
  mutation LoginUser($input: LoginInput!) {
    login(input: $input) {
      accessToken
      refreshToken
      user {
        id
        name
        email
      }
    }
  }
`;
