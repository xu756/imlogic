import {Outlet, useAccess} from '@umijs/max'

export default () => {
    const access = useAccess();
    return <Outlet/>;
}