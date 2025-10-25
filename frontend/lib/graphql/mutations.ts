import { gql } from "@apollo/client";

export const LOGIN_MUTATION = gql`
  mutation Login($input: LoginInput!) {
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

export const REGISTER_USER_MUTATION = gql`
  mutation RegisterUser($input: RegisterUserInput!) {
    registerUser(input: $input) {
      id
      name
      email
    }
  }
`;

export const RECOVER_PASSWORD_MUTATION = gql`
  mutation RecoverPassword($input: RecoverPasswordInput!) {
    recoverPassword(input: $input)
  }
`;

export const RECOVER_ACCOUNT_MUTATION = gql`
  mutation RecoverAccount($input: RecoverAccountInput!) {
    recoverAccount(input: $input)
  }
`;

export const VERIFY_EMAIL_MUTATION = gql`
  mutation VerifyEmail($input: VerifyEmailInput!) {
    verifyEmail(input: $input)
  }
`;

export const LOGOUT_MUTATION = gql`
  mutation Logout {
    logout
  }
`;

export const DELETE_ACCOUNT_MUTATION = gql`
  mutation DeleteAccount($input: DeleteAccountInput!) {
    deleteAccount(input: $input)
  }
`;
