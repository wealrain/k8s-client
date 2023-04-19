import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TablePagination from '@mui/material/TablePagination';
import Paper from '@mui/material/Paper';
import MoreVertBtn from '../MoreVertBtn';
import CircularProgress from '@mui/material/CircularProgress';
import Snackbar from '@mui/material/Snackbar';
import * as React from 'react';

const PAGE_SIZE = 10;

function ProgressMask(props) {
  const { loading } = props;
  return (
    <>
      {
        loading && 
        <>
          <div style={{ 
            position: 'absolute', 
            top: 0, 
            left: 0, 
            width: '100%', 
            height: '100%', 
            backgroundColor: 'rgba(255, 255, 255, 0.5)' }} />
          <CircularProgress style={{ position: 'absolute', top: '50%', left: '50%', transform: 'translate(-50%, -50%)' }} />
        </>
      }
    </>
  );
}

function PageinationTable(props) {
  const { 
      data, 
      columns,
      loading, 
      total, 
      current, 
      onChange,
      moreHandler,
      handleClick,
    } = props;

  const [snackbarState, setSnackbarState] = React.useState({
      open: false,
      vertical: 'top',
      horizontal: 'center',
    });
  const { vertical, horizontal, snackbarOpen } = snackbarState;
  const handleSnackbarClose = () => {
    setSnackbarState({ ...snackbarState, snackbarOpen: false });
  };
   
  return (
    <>
     <Snackbar
        anchorOrigin={{ vertical, horizontal }}
        open={snackbarOpen}
        onClose={handleSnackbarClose}
        message="todo"
        key={vertical + horizontal}
      />
        <TableContainer component={Paper} style={{position:'relative'}}>
          <ProgressMask loading={loading} />
          <Table sx={{ minWidth: 650 }} size="small" aria-label="a dense table">
            <TableHead>
              <TableRow>
                { columns.map((column,index) => (
                    <TableCell key={index}
                      sx={{ fontWeight: 'bold' }}
                      >{column.toUpperCase()}
                    </TableCell>
                ))}
                <TableCell> 
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {data.map((row,index) => (
                <TableRow
                  key={index}
                  sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                  {
                    columns.map((column,index) => (
                        <TableCell key={index} onClick={
                          (event) => {
                            if(handleClick) {
                              handleClick(event,row);
                            } else {
                              setSnackbarState({ ...snackbarState, snackbarOpen: true });
                            }
                          }
                        }
                        >{row[column]}</TableCell>
                    ))
                  }
                  <TableCell align='right' width={1}>
                    <MoreVertBtn items={moreHandler(row)} />
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        <TablePagination
        rowsPerPageOptions={[]}
        component="div"
        count={total}
        rowsPerPage={PAGE_SIZE}
        page={current}
        onPageChange={onChange}
      />
      </>
  );
}

export default PageinationTable;