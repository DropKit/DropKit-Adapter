import axios from "axios";
import { AxiosResponse } from "axios";
import colors from "colors";
import { sleep } from "sleep";

const SHOW_RESPONSE = true;

var CONFIG = require("../../configs/config.json");
const DROPKIT_ENDPOINT =
  "http://" + CONFIG.DROPKIT.HOST + ":" + CONFIG.DROPKIT.PORT;
const DROPKIT_ADMIN =
  "e0302f799ee2e6bcff366cf395dda8225e0a3ae9250740aeabf8e174a8d55c03";

const timestamp = () => {
  return new Date();
};

const NEEDLES_TABLE_NAME = `table${+timestamp()}`;

const DROPKIT_REQUEST = axios.create({
  baseURL: DROPKIT_ENDPOINT,
});

const DB_CREATE = async (data: any) => DROPKIT_REQUEST.post("/db/create", data);
const DB_INSERT = async (data: any) => DROPKIT_REQUEST.post("/db/insert", data);
const DB_UPDATE = async (data: any) => DROPKIT_REQUEST.post("/db/update", data);
const DB_SELECT = async (data: any) => DROPKIT_REQUEST.post("/db/select", data);

const PERMISSION_GRANT_OWNER = async (data: any) => {
  return await DROPKIT_REQUEST.post("/permission/grant/table/owner", data);
};
const PERMISSION_GRANT_MAINTAINER = async (data: any) => {
  return await DROPKIT_REQUEST.post("/permission/grant/table/maintainer", data);
};
const PERMISSION_GRANT_VIEWER = async (data: any) => {
  return await DROPKIT_REQUEST.post("/permission/grant/table/viewer", data);
};

const ROLE_CREATE = async (data: any) =>
  DROPKIT_REQUEST.post("/role/create", data);
const USER_CREATE = async () => DROPKIT_REQUEST.get("/user/create");

const ROLE_MANUFACTURER = ["id", "sn", "manufacturer", "mfd", "lc"];
const ROLE_LOGISTICS = ["id", "lc", "vehicle", "sign_date", "hospital"];
const ROLE_OPERATOR = ["id", "sign_date", "hospital", "shot_time", "patient"];
const ROLE_CONSUMER = [
  "id",
  "sn",
  "manufacturer",
  "mfd",
  "lc",
  "vehicle",
  "sign_date",
  "hospital",
  "shot_time",
  "patient",
];

const CREATE_TABLE = `CREATE TABLE ${NEEDLES_TABLE_NAME} (id int PRIMARY KEY, sn varchar, manufacturer varchar, mfd varchar, lc varchar, vehicle varchar, sign_date varchar, hospital varchar, shot_time varchar, patient varchar)`;

const RECORD_MANUFACTURER = [
  `INSERT INTO ${NEEDLES_TABLE_NAME} (id, sn, manufacturer, mfd) VALUES (`,
  `, '`,
  `', '`,
  `', '`,
  `')`,
];
const RECORD_LOGISTICS = [
  `UPDATE ${NEEDLES_TABLE_NAME} SET lc='`,
  `', vehicle='`,
  `', sign_date='`,
  `', hospital='`,
  `' WHERE id = `,
  ``,
];
const RECORD_OPERATOR = [
  `UPDATE ${NEEDLES_TABLE_NAME} SET shot_time='`,
  `', patient='`,
  `'  WHERE id = `,
  ``,
];

const READ_CONSUMER = `SELECT sn, manufacturer, mfd, lc, vehicle, sign_date, hospital, shot_time, patient FROM ${NEEDLES_TABLE_NAME} WHERE id = `;
const READ_VERIFIER = [
  `SELECT sn, manufacturer, mfd, lc, vehicle, sign_date, hospital, shot_time, patient FROM "${NEEDLES_TABLE_NAME}" WHERE id = 1`,
  `SELECT sn, manufacturer, mfd, lc, vehicle, sign_date, hospital, shot_time, patient FROM "${NEEDLES_TABLE_NAME}" WHERE id = 2`,
  `SELECT sn, manufacturer, mfd, lc, vehicle, sign_date, hospital, shot_time, patient FROM "${NEEDLES_TABLE_NAME}" WHERE id = 3`,
  `SELECT sn, manufacturer, mfd, lc, vehicle, sign_date, hospital, shot_time, patient FROM "${NEEDLES_TABLE_NAME}" WHERE id = 4`,
];

const sn = () => {
  const sn =
    Math.random().toString(36).substr(2, 3) +
    "-" +
    Math.random().toString(36).substr(2, 3) +
    "-" +
    Math.random().toString(36).substr(2, 4);
  return sn.toUpperCase();
};

let OWNER: string[] = [];
let MANUFACTURER: string[] = [];
let LOGISTICS: string[] = [];
let OPERATOR: string[] = [];
let CONSUMER: string[] = [];
let VERIFIER: string[] = [];

// create users
const step1 = async () => {
  let newUserInfo = await (await USER_CREATE()).data;
  OWNER.push(newUserInfo.PrivateKey);
  OWNER.push(newUserInfo.Account);

  newUserInfo = await (await USER_CREATE()).data;
  MANUFACTURER.push(newUserInfo.PrivateKey);
  MANUFACTURER.push(newUserInfo.Account);

  newUserInfo = await (await USER_CREATE()).data;
  LOGISTICS.push(newUserInfo.PrivateKey);
  LOGISTICS.push(newUserInfo.Account);

  newUserInfo = await (await USER_CREATE()).data;
  OPERATOR.push(newUserInfo.PrivateKey);
  OPERATOR.push(newUserInfo.Account);

  newUserInfo = await (await USER_CREATE()).data;
  CONSUMER.push(newUserInfo.PrivateKey);
  CONSUMER.push(newUserInfo.Account);

  newUserInfo = await (await USER_CREATE()).data;
  VERIFIER.push(newUserInfo.PrivateKey);
  VERIFIER.push(newUserInfo.Account);
};

// create table
const step2 = async () => {
  const res = await DB_CREATE({
    db_statement: CREATE_TABLE,
    caller_pk: DROPKIT_ADMIN,
  });
  return res.data;
  sleep(2);
};

// create role
const step3 = async () => {
  const res_permission = (
    await PERMISSION_GRANT_OWNER({
      user_name: OWNER[1],
      table_name: NEEDLES_TABLE_NAME,
      caller_pk: DROPKIT_ADMIN,
    })
  ).data;
  sleep(2);

  let res_role_create: any[] = [];
  res_role_create.push(
    (
      await ROLE_CREATE({
        columns: ROLE_MANUFACTURER,
        role_name: `${NEEDLES_TABLE_NAME}_ROLE_MANUFACTURER`,
        table_name: NEEDLES_TABLE_NAME,
        caller_pk: OWNER[0],
      })
    ).data
  );
  res_role_create.push(
    (
      await ROLE_CREATE({
        columns: ROLE_LOGISTICS,
        role_name: `${NEEDLES_TABLE_NAME}_ROLE_LOGISTICS`,
        table_name: NEEDLES_TABLE_NAME,
        caller_pk: OWNER[0],
      })
    ).data
  );
  res_role_create.push(
    (
      await ROLE_CREATE({
        columns: ROLE_OPERATOR,
        role_name: `${NEEDLES_TABLE_NAME}_ROLE_OPERATOR`,
        table_name: NEEDLES_TABLE_NAME,
        caller_pk: OWNER[0],
      })
    ).data
  );
  res_role_create.push(
    (
      await ROLE_CREATE({
        columns: ROLE_CONSUMER,
        role_name: `${NEEDLES_TABLE_NAME}_ROLE_CONSUMER`,
        table_name: NEEDLES_TABLE_NAME,
        caller_pk: OWNER[0],
      })
    ).data
  );
  res_role_create.push(
    (
      await ROLE_CREATE({
        columns: ROLE_CONSUMER,
        role_name: `${NEEDLES_TABLE_NAME}_ROLE_VERIFIER`,
        table_name: NEEDLES_TABLE_NAME,
        caller_pk: OWNER[0],
      })
    ).data
  );
  sleep(2);

  return [res_permission, res_role_create];
};

// grant permission
const step4 = async () => {
  let res: any[] = [];
  res.push(
    (
      await PERMISSION_GRANT_MAINTAINER({
        user_name: MANUFACTURER[1],
        column_role: `${NEEDLES_TABLE_NAME}_ROLE_MANUFACTURER`,
        table_name: NEEDLES_TABLE_NAME,
        caller_pk: OWNER[0],
      })
    ).data
  );
  res.push(
    (
      await PERMISSION_GRANT_MAINTAINER({
        user_name: LOGISTICS[1],
        column_role: `${NEEDLES_TABLE_NAME}_ROLE_LOGISTICS`,
        table_name: NEEDLES_TABLE_NAME,
        caller_pk: OWNER[0],
      })
    ).data
  );
  res.push(
    (
      await PERMISSION_GRANT_MAINTAINER({
        user_name: OPERATOR[1],
        column_role: `${NEEDLES_TABLE_NAME}_ROLE_OPERATOR`,
        table_name: NEEDLES_TABLE_NAME,
        caller_pk: OWNER[0],
      })
    ).data
  );
  sleep(3);
  return res;
};

// manufacturer record the data
const step5 = async () => {
  let res: any[] = [];
  for (let count = 1; count < 5; count++) {
    const db_statement = `${RECORD_MANUFACTURER[0]}${count}${
      RECORD_MANUFACTURER[1]
    }${sn()}${RECORD_MANUFACTURER[2]}MANUFACTURER_A${
      RECORD_MANUFACTURER[3]
    }${+timestamp()}${RECORD_MANUFACTURER[4]}`;
    res.push(
      (
        await DB_INSERT({
          db_statement: db_statement,
          caller_pk: MANUFACTURER[0],
        })
      ).data
    );
  }
  return res;
};

// logistics companies record the data
const step6 = async () => {
  let res: any[] = [];
  for (let count = 1; count < 5; count++) {
    const db_statement = `${RECORD_LOGISTICS[0]}Logistics_A${
      RECORD_LOGISTICS[1]
    }ABC-123${RECORD_LOGISTICS[2]}${+timestamp()}${
      RECORD_LOGISTICS[3]
    }Hospital_A${RECORD_LOGISTICS[4]}${count}`;
    res.push(
      (
        await DB_UPDATE({
          db_statement: db_statement,
          caller_pk: LOGISTICS[0],
        })
      ).data
    );
  }
  return res;
};

// hospital record the data
const step7 = async () => {
  let res: any[] = [];
  for (let count = 1; count < 5; count++) {
    const db_statement = `${RECORD_OPERATOR[0]}${+timestamp()}${
      RECORD_OPERATOR[1]
    }Patient_${count}${RECORD_OPERATOR[2]}${count}`;
    res.push(
      (
        await DB_UPDATE({
          db_statement: db_statement,
          caller_pk: OPERATOR[0],
        })
      ).data
    );
  }
  return res;
};

const step8 = async () => {
  const res = await PERMISSION_GRANT_VIEWER({
    user_name: CONSUMER[1],
    column_role: `${NEEDLES_TABLE_NAME}_ROLE_CONSUMER`,
    table_name: NEEDLES_TABLE_NAME,
    caller_pk: OWNER[0],
  });
  sleep(3);
  return res.data;
};

const step9 = async () => {
  let res: any[] = [];
  for (let count = 1; count < 5; count++) {
    const db_statement = `${READ_CONSUMER}${count}`;
    res.push(
      (
        await DB_SELECT({
          db_statement: db_statement,
          caller_pk: CONSUMER[0],
        })
      ).data
    );
  }
  return res;
};

const step10 = async () => {
  const db_statement = `${READ_CONSUMER}1`;
  const res = await DB_SELECT({
    db_statement: db_statement,
    caller_pk: MANUFACTURER[0],
  });
  return res.data;
};

const step11 = async () => {
  const db_statement = `${READ_CONSUMER}1`;
  const res = await DB_SELECT({
    db_statement: db_statement,
    caller_pk: LOGISTICS[0],
  });
  return res.data;
};

const step12 = async () => {
  const db_statement = `${READ_CONSUMER}1`;
  const res = await DB_SELECT({
    db_statement: db_statement,
    caller_pk: OPERATOR[0],
  });
  return res.data;
};

const step13 = async () => {
  const res = await PERMISSION_GRANT_VIEWER({
    user_name: VERIFIER[1],
    column_role: `${NEEDLES_TABLE_NAME}_ROLE_VERIFIER`,
    table_name: NEEDLES_TABLE_NAME,
    caller_pk: OWNER[0],
  });
  sleep(3);
  return res.data;
};

const step14 = async () => {
  let res: any[] = [];
  for (let count = 1; count < 5; count++) {
    const db_statement = `${READ_CONSUMER}${count}`;
    res.push(
      (
        await DB_SELECT({
          db_statement: db_statement,
          caller_pk: VERIFIER[0],
        })
      ).data
    );
  }
  return res;
};

jest.setTimeout(30000);
describe("Regression Test", () => {
  test("step2:", async () => {
    await step1();
    const response = await step2();
    expect(response.Code).toBe(0);
    expect(response.Message).toBe("Ok");
    expect(response.Audit).toMatch(new RegExp(/0x[a-z0-9]{64}/));
  });

  test("step3", async () => {
    const [res_permission, res_role_create] = await step3();
    expect(res_permission.Code).toBe(0);
    expect(res_permission.Message).toBe("Ok");

    if (Array.isArray(res_role_create)) {
      res_role_create.forEach((element) => {
        expect(element.Code).toBe(0);
        expect(element.Message).toBe("Ok");
      });
    }
  });

  test("step4", async () => {
    const res = await step4();
    if (Array.isArray(res)) {
      res.forEach((element) => {
        expect(element.Code).toBe(0);
        expect(element.Message).toBe("Ok");
      });
    }
  });

  test("step5", async () => {
    const res = await step5();
    res.forEach((element) => {
      expect(element.Code).toBe(0);
      expect(element.Message).toBe("Ok");
      expect(element.Audit).toMatch(new RegExp(/0x[a-z0-9]{64}/));
    });
  });
  test("step6", async () => {
    const res = await step6();
    res.forEach((element) => {
      expect(element.Code).toBe(0);
      expect(element.Message).toBe("Ok");
      expect(element.Audit).toMatch(new RegExp(/0x[a-z0-9]{64}/));
    });
  });
  test("step7", async () => {
    const res = await step7();
    res.forEach((element) => {
      expect(element.Code).toBe(0);
      expect(element.Message).toBe("Ok");
      expect(element.Audit).toMatch(new RegExp(/0x[a-z0-9]{64}/));
    });
  });
  test("step8", async () => {
    const response = await step8();
    expect(response.Code).toBe(0);
    expect(response.Message).toBe("Ok");
  });
  test("step9", async () => {
    const res = await step9();
    res.forEach((element) => {
      expect(element.Code).toBe(0);
      expect(element.Message).toBe("Ok");
      expect(element.Audit).toMatch(new RegExp(/0x[a-z0-9]{64}/));
    });
  });
  test("step10", async () => {
    const response = await step10();
    expect(response.Code).toBe(0);
    expect(response.Message).toBe("Ok");
    expect(response.Audit).toMatch(new RegExp(/0x[a-z0-9]{64}/));
  });
  test("step11", async () => {
    const response = await step11();
    expect(response.Code).toBe(0);
    expect(response.Message).toBe("Ok");
    expect(response.Audit).toMatch(new RegExp(/0x[a-z0-9]{64}/));
  });
  test("step12", async () => {
    const response = await step12();
    expect(response.Code).toBe(0);
    expect(response.Message).toBe("Ok");
    expect(response.Audit).toMatch(new RegExp(/0x[a-z0-9]{64}/));
  });
  test("step13", async () => {
    const response = await step13();
    expect(response.Code).toBe(0);
    expect(response.Message).toBe("Ok");
  });
  test("step14", async () => {
    const res = await step14();
    res.forEach((element) => {
      expect(element.Code).toBe(0);
      expect(element.Message).toBe("Ok");
    });
  });
});
