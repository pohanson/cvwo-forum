interface IUser {
  username: string;
  name: string;
  role: number;
}
export class User {
  public username: string;
  public name: string;
  public role: Role;

  private constructor(username: string, name: string, role: Role) {
    this.username = username;
    this.name = name;
    this.role = role;
  }

  public static fromJsonStr(jsonStr: string): User {
    const obj = JSON.parse(jsonStr) as IUser;
    const requiredField = ["username", "name", "role"];
    if (requiredField.some((val) => !Object.hasOwn(obj, val))) {
      throw new TypeError(
        `User json missing required fields. Received: ${jsonStr}`
      );
    }
    if (obj.role > 3 || obj.role < 1) {
      throw new RangeError(`Invalid json given. Received ${jsonStr}`);
    }

    return new User(obj.username, obj.name, obj.role);
  }
  public toJsonStr(): string {
    return JSON.stringify({
      username: this.username,
      name: this.name,
      role: this.role,
    });
  }
}
enum Role {
  "normal" = 1,
  "admin" = 2,
  "moderator" = 3,
}
