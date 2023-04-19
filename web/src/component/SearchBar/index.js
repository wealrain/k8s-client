import { useLocation } from 'react-router-dom';
import {Box ,Breadcrumbs,Paper,IconButton,InputBase,} from '@mui/material'
import {Search as SearchIcon} from '@mui/icons-material'
import NavigateNextIcon from '@mui/icons-material/NavigateNext';
import * as React from 'react';

export default function SearchBar(props) {

    const { onSearchResource } = props;

    const [searchValue, setSearchValue] = React.useState('');

    const location = useLocation();
    const pathnames = location.pathname.split('/').filter((x) => x);
    const breadcrumbs = pathnames.map((value, index) => {
        const last = index === pathnames.length - 1;
        const to = `/${pathnames.slice(0, index + 1).join('/')}`;
        return last ? (
            <Box key={to} sx={{color:"#000000"}}>{value}</Box>
        ) : (
            <Box key={to} sx={{color:"#0000008a"}}>{value}</Box>
        );
    });
    return (
        <Box sx={{ display:"flex",mb:2}}>
             <Breadcrumbs
                separator={<NavigateNextIcon fontSize="small" />}
                sx={{flexGrow:1}}
            >
            {breadcrumbs}
            </Breadcrumbs>
    <Paper
      component="form"
      sx={{ p: '2px 4px', display: 'flex', alignItems: 'center', width: 400 }}
      size='small'
    >
     
      <InputBase
        sx={{ ml: 1, flex: 1 }}
        placeholder="Search Resource By Name"
        size='small'
        value={searchValue}
        onChange={(e) => {
            setSearchValue(e.target.value)
        }}
        onKeyDown={(e) => {
            if (e.key === 'Enter') {
                onSearchResource(e.target.value)
            }
        }}    
      />
      <IconButton type="button" sx={{ p: '10px' }} 
        onClick={() => {
            onSearchResource(searchValue)
        }}
      >
        <SearchIcon />
      </IconButton>
     
    </Paper>
        </Box>
    )
}