# Write your MySQL query statement below

select dp.name as `Department`, em.name as `Employee` , em.salary as `Salary`
    from Employee em 
        left join Department as dp on em.departmentId = dp.id
    where
    3 > ( select count(distinct em2.Salary)
        from Employee em2
        where
            em2.Salary > em.Salary
            and em.DepartmentId = em2.DepartmentId
        )