import {PageContainer} from '@ant-design/pro-components';
import {useModel} from '@umijs/max';
import styles from './index.scss';

const HomePage: React.FC = () => {
    const {config,} = useModel('global');
    return (
        <PageContainer ghost>
            <div className={styles.container}>
                {config.title}
            </div>
        </PageContainer>
    );
};

export default HomePage;
