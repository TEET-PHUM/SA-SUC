import React, { useState, FC, useEffect, Fragment } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import Avatar from '@material-ui/core/Avatar';
import { deepOrange } from '@material-ui/core/colors';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';

import { EntUser } from '../../api/models/EntUser'; // import interface User
import { EntTaste } from '../../api/models/EntTaste'; // import interface Taste
import { EntMealplan } from '../../api/models/EntMealplan'; // import interface Mealplan
import { EntFoodmenu } from '../../api/models/EntFoodmenu'; // import interface Foodmenu
import EatinghistoryTable from '../Table';

 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
  root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
  },
  margin: {
    margin: theme.spacing(3),
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: '25ch',
  },
  orange: { 
    color: theme.palette.getContrastText(deepOrange[500]),
    backgroundColor: deepOrange[500],
  },
}),);
 
export default function EatingHistory() {
  const classes = useStyles();
  const api = new DefaultApi();

  const [users, setUsers] = useState<EntUser[]>([]);
  const [foodmenus, setFoodmenus] = useState<EntFoodmenu[]>([]);
  const [tastes, setTastes] = useState<EntTaste[]>([]);
  const [mealplans, setMealplans] = useState<EntMealplan[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [addedtime, setAddedtime] = useState(String);
  const [mealplanid, setMealplan] = useState(Number);
  const [foodmenuid, setFoodmenu] = useState(Number);
  const [tasteid, setTaste] = useState(Number);
  const [userid, setUser] = useState(Number);


  useEffect(() => {
    const getUsers = async () => {
      const res = await api.listUser({ limit: 4, offset: 0 });
      setLoading(false);
      setUsers(res);
      console.log(res);
    };
    getUsers();

    const getFoodmenu = async () => {
      const res = await api.listFoodmenu({ limit: 6, offset: 0 });
      setLoading(false);
      setFoodmenus(res);
      console.log(res);
    };
    getFoodmenu();

    const getTaste = async () => {
      const res = await api.listTaste({ limit: 5, offset: 0 });
      setLoading(false);
      setTastes(res);
      console.log(res);
    };
    getTaste();

    const getMealplan = async () => {
      const res = await api.listMealplan({ limit: 6, offset: 0 });
      setLoading(false);
      setMealplans(res);
      console.log(res);
    };
    getMealplan();

  }, [loading]);

  const handleDatetimeChange = (event: any) => {
    setAddedtime(event.target.value as string);
  };

  const CreateEatinghistory = async () => {
    const eatinghistory = {
       addedTime     : addedtime + ":00+07:00",
       foodmenuID    : foodmenuid,
       mealplanID    : mealplanid,
       userID        : userid,
       tasteID       : tasteid,
    }
    console.log(eatinghistory);
    const res:any = await api.createEatinghistory({ eatinghistory : eatinghistory});
      setStatus(true);
        if (res.id != ''){
      setAlert(true);
        } else {
      setAlert(false);
        }

    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);

  };

  const useridhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };

  const tasteidhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setTaste(event.target.value as number);
  };

  const foodmenuidhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setFoodmenu(event.target.value as number);
  };
  
  const mealplanidhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setMealplan(event.target.value as number);
  };
  

  return (
    <Page theme={pageTheme.home}>
      <Header
        title="ระบบบันทึกการรับประทานอาหาร"
        subtitle="Subsystem to record meals or food to Eatinghistory"
        >
         <Avatar className={classes.orange}>U</Avatar>  
         
       </Header>
      <Content>
        <ContentHeader title="เพิ่มประวัติการกิน">
        <Link component={RouterLink} to="/eateat">
         <Button style={{ marginLeft: 20 }} variant="contained" color="secondary" >
            ประวัติ
         </Button>
        </Link>
        {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 Nice!!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 Error!!
               </Alert>
             )}
           </div>
         ) : null}
        </ContentHeader>
        <Container maxWidth="sm">
           <Grid container spacing={3}>
             <Grid item xs={12}></Grid>
             <Grid item xs={3}>
               <div className={classes.paper}>มื้ออาหารที่จัดไว้</div>
             </Grid>
             <Grid item xs={9}>
               <FormControl variant="outlined" className={classes.formControl}>
                 <InputLabel id="mealplan_id-label">เลือก Meal_plan</InputLabel>
                <Select
                   labelId="mealplan_id-label"
                   label="Mealplan"
                   id="mealplan_id"
                   value={mealplanid}
                   onChange={mealplanidhandleChange}
                   style = {{width: 600}}
                >
                  {mealplans.map((item:EntMealplan)=>
                    <MenuItem value={item.id}>{item.mealplanName}</MenuItem>)}
                </Select>
               </FormControl>
             </Grid>
 
             <Grid item xs={3}>
               <div className={classes.paper}>รายการอาหาร</div>
             </Grid>
             <Grid item xs={9}>
               <FormControl variant="outlined" className={classes.formControl}>
                 <InputLabel id="foodmenu_id-label">เลือก Food_menu</InputLabel>
                <Select
                  labelId="foodmenu_id-label"
                  label="Foodmenu"
                  id="foodmenu_id"
                  value={foodmenuid}
                  onChange={foodmenuidhandleChange}
                  style = {{width: 600}}
               >
                 {foodmenus.map((item:EntFoodmenu)=>
                   <MenuItem value={item.id}>{item.foodmenuName}</MenuItem>)}
                  </Select>
               </FormControl>
             </Grid>
 
             <Grid item xs={3}>
               <div className={classes.paper}>ความพึงพอใจ</div>
             </Grid>
             <Grid item xs={9}>
               <FormControl variant="outlined" className={classes.formControl}>
                 <InputLabel id="taste_id-label">เลือก Taste</InputLabel>
                <Select
                  labelId="taste_id-label"
                  label="Taste"
                  id="taste_id"
                  value={tasteid}
                  onChange={tasteidhandleChange}
                  style = {{width: 600}}
               >
                 {tastes.map((item:EntTaste)=>
                   <MenuItem value={item.id}>{item.tasteName}</MenuItem>)}
                  </Select>
               </FormControl>
             </Grid>
 
             <Grid item xs={3}>
               <div className={classes.paper}>ผู้ใช้</div>
             </Grid>
             <Grid item xs={9}>
               <FormControl variant="outlined" className={classes.formControl}>
                 <InputLabel id="user_id-label">เลือก User</InputLabel>
                <Select
                  labelId="user_id-label"
                  label="User"
                  id="user_id"
                  value={userid}
                  onChange={useridhandleChange}
                  style = {{width: 600}}
               >
                 {users.map((item:EntUser)=>
                   <MenuItem value={item.id}>{item.email}</MenuItem>)}
                  </Select>
               </FormControl>
             </Grid>
 
             <Grid item xs={3}>
               <div className={classes.paper} >เวลาที่รับประทาน</div>
             </Grid>
             <Grid item xs={9}>
               <div>
               <FormControl
                className={classes.margin}
                variant="outlined"
              >
               <TextField
                   id="datetime"
                   label="DateTime"
                   type="datetime-local"
                   value={addedtime}
                   onChange={handleDatetimeChange}
                   className={classes.textField}
                   InputLabelProps={{
                     shrink: true,
                   }}
                />
               </FormControl>
               </div>
             </Grid>
 
             <Grid item xs={3}></Grid>
             <Grid item xs={9}>
               <Button
                 variant="contained"
                 color="primary"
                 size="large"
                 style={{ marginRight: 1 }}
                 onClick={() => {
                  CreateEatinghistory();
                }}
               >
                 บันทึก
               </Button>
             </Grid>
           </Grid>
         </Container>
      </Content>
    </Page>
  );
};