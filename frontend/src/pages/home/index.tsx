import {PageContainer} from '@ant-design/pro-components';
import {useModel} from '@umijs/max';
import styles from './index.scss';

const HomePage: React.FC = () => {
    const global = useModel('global');
    console.log(global);
    return (
        <PageContainer ghost>
            <div className={styles.container}>
                <h1 className={styles.title}>Welcome to Max!</h1>
            </div>
        </PageContainer>
    );
};

export default HomePage;
